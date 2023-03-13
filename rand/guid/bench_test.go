package bench_test

import (
	"testing"

	"github.com/bsm/nanoid"
	"github.com/chilts/sid"
	"github.com/fogfish/guid/v2"
	"github.com/google/uuid"
	"github.com/itrabbit/rid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid/v3"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"github.com/wmentor/uniq"
)

func uuidNew() string     { return uuid.New().String() }
func gonanoidNew() string { s, _ := gonanoid.Nanoid(); return s }
func guidNew() string     { return guid.G(guid.Clock).String() }
func ridNew() string      { return rid.New().String() }
func xidNew() string      { return xid.New().String() }
func ksuidNew() string    { return ksuid.New().String() }

func bench(b *testing.B, fn func() string, min, max int) {
	for i := 0; i < b.N; i++ {
		if s := fn(); len(s) < min || len(s) > max {
			b.Fatalf("expected %d-%d chars, but got %d (%s)", min, max, len(s), s)
		}
	}
}

func pbench(b *testing.B, fn func() string, min, max int) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if s := fn(); len(s) < min || len(s) > max {
				b.Fatalf("expected %d-%d chars, but got %d (%s)", min, max, len(s), s)
			}
		}
	})
}

func Benchmark_random(b *testing.B) {
	b.Run("bsm/nanoid", func(b *testing.B) { // "Q2mAgYI9yIgSPd9ENknGM"
		bench(b, nanoid.New, 21, 21)
	})
	b.Run("google/uuid", func(b *testing.B) { // "6887c36e-85ba-4d64-80fe-0e14ee38e2a2"
		bench(b, uuidNew, 36, 36)
	})
	b.Run("lithammer/shortuuid", func(b *testing.B) { // "DSNxRpoZZPa6Eb9GkMD"
		bench(b, shortuuid.New, 18, 22)
	})
	b.Run("matoous/go-nanoid", func(b *testing.B) { // "rhANQ2U-QMnfNVvDnZaSs"
		bench(b, gonanoidNew, 21, 21)
	})
}

func Benchmark_random_parallel(b *testing.B) {
	b.Run("bsm/nanoid", func(b *testing.B) { // "Q2mAgYI9yIgSPd9ENknGM"
		pbench(b, nanoid.New, 21, 21)
	})
	b.Run("google/uuid", func(b *testing.B) { // "6887c36e-85ba-4d64-80fe-0e14ee38e2a2"
		pbench(b, uuidNew, 36, 36)
	})
	b.Run("lithammer/shortuuid", func(b *testing.B) { // "DSNxRpoZZPa6Eb9GkMD"
		pbench(b, shortuuid.New, 18, 22)
	})
	b.Run("matoous/go-nanoid", func(b *testing.B) { // "rhANQ2U-QMnfNVvDnZaSs"
		pbench(b, gonanoidNew, 21, 21)
	})
}

func Benchmark_sorted(b *testing.B) {
	b.Run("chilts/sid", func(b *testing.B) { // "1599734701420031946-7468365110993091186"
		bench(b, sid.Id, 39, 39)
	})
	b.Run("fogfish/guid", func(b *testing.B) { // "NgOglOfNRUy5c7.0"
		bench(b, guidNew, 16, 16)
	})
	b.Run("itrabbit/rid", func(b *testing.B) { // "8n2q609dtz97m42o67e0"
		bench(b, ridNew, 20, 20)
	})
	b.Run("kjk/betterguid", func(b *testing.B) { // "-MGrcC-au8zVbV_nR7mH"
		bench(b, betterguid.New, 20, 20)
	})
	b.Run("rs/xid", func(b *testing.B) { // "btd08je7gngoc3t23on0"
		bench(b, xidNew, 20, 20)
	})
	b.Run("segmentio/ksuid", func(b *testing.B) { // "1hJgyJ1FomiECVaJTV1kWmbzoQE"
		bench(b, ksuidNew, 27, 27)
	})
	b.Run("wmentor/uniq", func(b *testing.B) { // "a7b1d7c7d1458316336661f00320551368"
		bench(b, uniq.New, 32, 34)
	})
}

func Benchmark_sorted_parallel(b *testing.B) {
	b.Run("chilts/sid", func(b *testing.B) { // "1599734701420031946-7468365110993091186"
		pbench(b, sid.Id, 39, 39)
	})
	b.Run("fogfish/guid", func(b *testing.B) { // "NgOglOfNRUy5c7.0"
		pbench(b, guidNew, 16, 16)
	})
	b.Run("itrabbit/rid", func(b *testing.B) { // "8n2q609dtz97m42o67e0"
		pbench(b, ridNew, 20, 20)
	})
	b.Run("kjk/betterguid", func(b *testing.B) { // "-MGrcC-au8zVbV_nR7mH"
		pbench(b, betterguid.New, 20, 20)
	})
	b.Run("rs/xid", func(b *testing.B) { // "btd08je7gngoc3t23on0"
		pbench(b, xidNew, 20, 20)
	})
	b.Run("segmentio/ksuid", func(b *testing.B) { // "1hJgyJ1FomiECVaJTV1kWmbzoQE"
		pbench(b, ksuidNew, 27, 27)
	})
	b.Run("wmentor/uniq", func(b *testing.B) { // "a7b1d7c7d1458316336661f00320551368"
		pbench(b, uniq.New, 32, 34)
	})
}
