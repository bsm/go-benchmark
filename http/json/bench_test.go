package bench_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/valyala/fasthttp"
)

var (
	small = struct {
		Code   int
		Status string
	}{Code: 200, Status: "OK"}

	large = struct {
		Strs []string
		Nums []int
	}{}
)

func init() {
	for i := 0; i < 64; i++ {
		large.Strs = append(large.Strs, strings.Repeat("x", 512))
	}
	for i := 0; i < 1024; i++ {
		large.Nums = append(large.Nums, i*2)
	}
}

// --------------------------------------------------------------------

func Benchmark(b *testing.B) {
	b.Run("small payload", func(b *testing.B) {
		benchmark(b, 27, &small)
	})
	b.Run("large payload", func(b *testing.B) {
		benchmark(b, 37545, &large)
	})
}

func benchmark(b *testing.B, contentLen int, v interface{}) {
	b.Run("direct", func(b *testing.B) {
		benchServe(b, contentLen, func(w http.ResponseWriter, r *http.Request) {
			if err := json.NewEncoder(w).Encode(v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})

	b.Run("bufio", func(b *testing.B) {
		benchServe(b, contentLen, func(w http.ResponseWriter, r *http.Request) {
			bw := bufio.NewWriter(w)
			if err := json.NewEncoder(bw).Encode(v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := bw.Flush(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})

	b.Run("bufio pool", func(b *testing.B) {
		pool := sync.Pool{
			New: func() interface{} { return bufio.NewWriter(nil) },
		}

		benchServe(b, contentLen, func(w http.ResponseWriter, r *http.Request) {
			bw := pool.Get().(*bufio.Writer)
			bw.Reset(w)
			defer pool.Put(bw)

			if err := json.NewEncoder(bw).Encode(v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := bw.Flush(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})

	b.Run("buffer", func(b *testing.B) {
		benchServe(b, contentLen, func(w http.ResponseWriter, r *http.Request) {
			bb := new(bytes.Buffer)
			if err := json.NewEncoder(bb).Encode(v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if _, err := bb.WriteTo(w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})

	b.Run("buffer pool", func(b *testing.B) {
		pool := sync.Pool{
			New: func() interface{} { return new(bytes.Buffer) },
		}

		benchServe(b, contentLen, func(w http.ResponseWriter, r *http.Request) {
			bb := pool.Get().(*bytes.Buffer)
			bb.Reset()
			defer pool.Put(bb)

			if err := json.NewEncoder(bb).Encode(v); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if _, err := bb.WriteTo(w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})
}

func benchServe(b *testing.B, contentLen int, fn http.HandlerFunc) {
	srv := httptest.NewServer(http.HandlerFunc(fn))
	defer srv.Close()

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		req.Reset()
		req.SetRequestURI(srv.URL)
		res.Reset()

		err := fasthttp.Do(req, res)
		if err != nil {
			b.Fatalf("expected no error, but received %v", err)
		} else if code := res.StatusCode(); code != 200 {
			b.Fatalf("expected 200 response but received %v", code)
		} else if n := len(res.Body()); n != contentLen {
			b.Fatalf("expected content length to be %d, but was %d", contentLen, n)
		}

		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}
}
