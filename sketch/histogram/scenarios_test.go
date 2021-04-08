package histogram_test

import (
	"math"
	"math/rand"
	"sort"
)

var scenarios = []*Scenario{
	{
		Name: "uniform",
		Next: func(r *rand.Rand) float64 { return r.Float64() * 200 },
	},
	{
		Name: "lognorm",
		Next: func(r *rand.Rand) float64 {
			return 100 * math.Exp(r.NormFloat64()*0.9)
		},
	},
}

const numSeeds = 100_000

// ----------------------------------------------------------------------------

func init() {
	rnd := rand.New(rand.NewSource(1))

	for _, scn := range scenarios {
		scn.init(rnd)
	}
}

// Scenario is a benchmark scenario.
type Scenario struct {
	Name       string
	Next       func(*rand.Rand) float64
	Data       []float64
	Candidates []Sketch
}

func (s *Scenario) init(r *rand.Rand) {
	for _, new := range factories {
		s.Candidates = append(s.Candidates, new())
	}
	s.Data = make([]float64, numSeeds)
	for i := range s.Data {
		x := s.Next(r)
		s.Data[i] = x

		for _, cnd := range s.Candidates {
			cnd.Add(x)
		}
	}
	sort.Float64s(s.Data)
}
