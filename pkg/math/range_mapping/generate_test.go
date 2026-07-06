package range_mapping

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestGenerate(t *testing.T) {
	assert.Any(
		t,
		[]*Mapping{
			{Range: ranges.Range{L: 0, R: 0.625}, Value: "Alfa"},
			{Range: ranges.Range{L: 0.625, R: 1.25}, Value: "Bravo"},
			{Range: ranges.Range{L: 1.25, R: 1.875}, Value: "Charlie"},
			{Range: ranges.Range{L: 1.875, R: 2.5}, Value: "Delta"},
		},
		Generate(
			[]float64{0.5, 1.5, 2.5},
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
	assert.Any(
		t,
		[]*Mapping{
			{
				Range: ranges.Range{
					L: 0,
					R: 127673.40987711631,
				},
				Value: "Alfa",
			},
			{
				Range: ranges.Range{
					L: 127673.40987711631,
					R: 255346.81975423262,
				},
				Value: "Bravo",
			},
			{
				Range: ranges.Range{
					L: 255346.81975423262,
					R: 383020.2296313489,
				},
				Value: "Charlie",
			},
			{
				Range: ranges.Range{
					L: 383020.2296313489,
					R: 510693.63950846525,
				},
				Value: "Delta",
			},
		},
		Generate(
			[]float64{
				2218.642533079671,
				12796.918271387178,
				4732.263995009753,
				38236.105391612495,
				75860.42701528987,
				7971.654963137754,
				8560.167475137754,
				38809.12563404187,
				59919.87512489732,
				83477.34142626717,
				510693.63950846525,
				2037.8545633692056,
				35611.377112848655,
				53659.07512490345,
				72744.67293312878,
				546.6313734575342,
				1573.4872086356165,
				226935.19567286095,
				52517.80433039517,
				2234.058753189699,
				54119.91983725063,
				141070.4662482095,
				17096.431646127345,
				93383.2473989006,
				85924.20082355813,
				179770.0257002705,
				38682.81088646225,
				5370.08292679715,
				78462.7853715095,
			},
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
}
