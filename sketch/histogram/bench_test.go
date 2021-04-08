package histogram_test

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	// use uniform data
	data := scenarios[0].Data

	for _, new := range factories {
		sketch := new()
		b.Run(sketch.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sketch.Add(data[i%len(data)])
			}
		})
	}
}

func BenchmarkQuery(b *testing.B) {
	for _, scn := range scenarios {
		b.Run(scn.Name, func(b *testing.B) {
			for _, sketch := range scn.Candidates {
				b.Run(sketch.Name(), func(b *testing.B) {
					for i := 0; i < b.N; i += 4 {
						sketch.Query(0.1)
						sketch.Query(0.5)
						sketch.Query(0.9)
						sketch.Query(0.99)
					}
				})
			}
		})
	}
}

func BenchmarkPrecision(b *testing.B) {
	percentiles := []float64{.1, .25, .5, .75, .9, .95, .99}

	for _, scn := range scenarios {
		b.Run(scn.Name, func(b *testing.B) {
			for _, pct := range percentiles {
				exp := scn.Data[int(pct*float64(len(scn.Data)))-1] // actual quantile

				b.Run(fmt.Sprintf("p%.0f", pct*100), func(b *testing.B) {
					for _, sketch := range scn.Candidates {
						b.Run(sketch.Name(), func(b *testing.B) {
							got := sketch.Query(pct)
							rel := math.Abs(got-exp) / (math.Abs(got) + math.Abs(exp)) * 200
							b.ReportMetric(0, "ns/op")
							b.ReportMetric(rel, "%err")
						})
					}
				})
			}
		})
	}
}
