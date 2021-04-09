package histogram_test

import (
	ajwerner "github.com/ajwerner/tdigest"
	beorn7 "github.com/beorn7/perks/quantile"
	bsmhisto "github.com/bsm/histogram"
	bsmdigest "github.com/bsm/tdigest"
	caio "github.com/caio/go-tdigest"
	influx "github.com/influxdata/tdigest"
	segmentio "github.com/segmentio/tdigest"
	signalfx "github.com/signalfx/tdigest"
	spenczar "github.com/spenczar/tdigest"
	willbeason "github.com/willbeason/tdigest/pkg/tdigest"
	// xavier268 "github.com/xavier268/tdigest" - too slow for this test
)

var factories = []func() Sketch{
	func() Sketch { return ajwernerSN{ajwerner.New()} },
	func() Sketch { return beorn7SN{beorn7.NewHighBiased(1e-2)} },
	func() Sketch { return bsmhistoSN{bsmhisto.New(100)} },
	func() Sketch { return bsmdigestSN{bsmdigest.NewWithCompression(100)} },
	func() Sketch {
		td, _ := caio.New(caio.Compression(100))
		return caioSN{td}
	},
	func() Sketch { return influxSN{influx.NewWithCompression(100)} },
	func() Sketch { return spenczarSN{spenczar.NewWithCompression(100)} },
	func() Sketch { return segmentioSN{segmentio.NewWithCompression(100)} },
	func() Sketch { return signalfxSN{signalfx.NewWithCompression(100)} },
	func() Sketch { return willbeasonSN{willbeason.New(100)} },
}

// Sketch abstraction.
type Sketch interface {
	Name() string
	Add(float64)
	Query(float64) float64
}

// ----------------------------------------------------------------------------

type ajwernerSN struct{ *ajwerner.TDigest }

func (s ajwernerSN) Name() string            { return "ajwerner" }
func (s ajwernerSN) Add(x float64)           { s.Record(x) }
func (s ajwernerSN) Query(p float64) float64 { return s.ValueAt(p) }

type beorn7SN struct{ *beorn7.Stream }

func (s beorn7SN) Name() string  { return "beorn7" }
func (s beorn7SN) Add(x float64) { s.Insert(x) }

type bsmhistoSN struct{ *bsmhisto.Histogram }

func (s bsmhistoSN) Name() string            { return "bsmhisto" }
func (s bsmhistoSN) Query(p float64) float64 { return s.Quantile(p) }

type bsmdigestSN struct{ *bsmdigest.TDigest }

func (s bsmdigestSN) Name() string            { return "bsmdigest" }
func (s bsmdigestSN) Add(x float64)           { s.TDigest.Add(x, 1) }
func (s bsmdigestSN) Query(p float64) float64 { return s.Quantile(p) }

type caioSN struct{ *caio.TDigest }

func (s caioSN) Name() string            { return "caio" }
func (s caioSN) Add(x float64)           { _ = s.TDigest.Add(x) }
func (s caioSN) Query(p float64) float64 { return s.Quantile(p) }

type influxSN struct{ *influx.TDigest }

func (s influxSN) Name() string            { return "influx" }
func (s influxSN) Add(x float64)           { s.TDigest.Add(x, 1) }
func (s influxSN) Query(p float64) float64 { return s.Quantile(p) }

type spenczarSN struct{ *spenczar.TDigest }

func (s spenczarSN) Name() string            { return "spenczar" }
func (s spenczarSN) Add(x float64)           { s.TDigest.Add(x, 1) }
func (s spenczarSN) Query(p float64) float64 { return s.Quantile(p) }

type segmentioSN struct{ *segmentio.TDigest }

func (s segmentioSN) Name() string            { return "segmentio" }
func (s segmentioSN) Add(x float64)           { s.TDigest.Add(x, 1) }
func (s segmentioSN) Query(p float64) float64 { return s.Quantile(p) }

type signalfxSN struct{ *signalfx.TDigest }

func (s signalfxSN) Name() string            { return "signalfx" }
func (s signalfxSN) Add(x float64)           { s.TDigest.Add(x, 1) }
func (s signalfxSN) Query(p float64) float64 { return s.Quantile(p) }

type willbeasonSN struct{ *willbeason.TDigest }

func (s willbeasonSN) Name() string            { return "willbeason" }
func (s willbeasonSN) Query(p float64) float64 { return s.Quantile(p) }
