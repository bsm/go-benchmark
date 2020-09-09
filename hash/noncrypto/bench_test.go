package bench_test

import (
	"hash"
	"hash/crc64"
	"hash/fnv"
	"math/rand"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/creachadair/cityhash"
	"github.com/dgryski/go-farm"
	"github.com/dgryski/go-metro"
	"github.com/dgryski/go-spooky"
	"github.com/dgryski/go-stadtx"
	"github.com/dgryski/go-wyhash"
	"github.com/mmcloughlin/meow"
	"github.com/shivakar/metrohash"
	"github.com/spaolacci/murmur3"
)

var (
	p64   = make([]byte, 64)
	p256  = make([]byte, 256)
	p4096 = make([]byte, 4096)
)

// seed fixtures
func init() {
	rnd := rand.New(rand.NewSource(33))
	if _, err := rnd.Read(p64); err != nil {
		panic(err)
	}
	if _, err := rnd.Read(p256); err != nil {
		panic(err)
	}
	if _, err := rnd.Read(p4096); err != nil {
		panic(err)
	}
}

// --------------------------------------------------------------------

func Benchmark_64byte(b *testing.B) {
	benchmark(b, p64)
}

func Benchmark_256byte(b *testing.B) {
	benchmark(b, p256)
}

func Benchmark_4096byte(b *testing.B) {
	benchmark(b, p4096)
}

func benchmark(b *testing.B, p []byte) {
	b.Run("crc64", func(b *testing.B) {
		h := crc64.New(crc64.MakeTable(crc64.ISO))
		benchHash(b, h, p)
	})
	b.Run("fnv1", func(b *testing.B) {
		h := fnv.New64()
		benchHash(b, h, p)
	})
	b.Run("cespare/xxhash", func(b *testing.B) {
		h := xxhash.New()
		benchHash(b, h, p)
	})
	b.Run("creachadair/cityhash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cityhash.Hash64(p)
		}
	})
	b.Run("dgryski/go-farm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			farm.Hash64(p)
		}
	})
	b.Run("dgryski/go-metro", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			metro.Hash64(p, 0)
		}
	})
	b.Run("dgryski/go-spooky", func(b *testing.B) {
		h := spooky.New(1, 2)
		benchHash(b, h, p)
	})
	b.Run("dgryski/go-stadtx", func(b *testing.B) {
		state := stadtx.SeedState([]uint64{1, 2})
		for i := 0; i < b.N; i++ {
			stadtx.Hash(&state, p)
		}
	})
	b.Run("dgryski/go-wyhash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			wyhash.Hash(p, 0)
		}
	})
	b.Run("mmcloughlin/meow", func(b *testing.B) {
		h := meow.New64(0)
		benchHash(b, h, p)
	})
	b.Run("shivakar/metrohash", func(b *testing.B) {
		h := metrohash.NewMetroHash64()
		benchHash(b, h, p)
	})
	b.Run("spaolacci/murmur3", func(b *testing.B) {
		h := murmur3.New64()
		benchHash(b, h, p)
	})
}

func benchHash(b *testing.B, h hash.Hash64, p []byte) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(p)
		h.Sum64()
	}
}
