package range_mapping

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"sort"
)

func Generate(
	f []float64,
	values []string,
) []*Mapping {
	result := make([]*Mapping, 0)

	if len(f) == 0 || len(values) == 0 {
		return result
	}

	sort.Float64s(f)
	rangeCount := len(values)

	if rangeCount == 0 {
		return result
	}

	maximum := f[len(f)-1]
	perRange := maximum / float64(rangeCount)

	for i := 0; i < rangeCount; i++ {
		m := New()
		m.Range = ranges.Range{
			L: float64(i) * perRange,
			R: float64(i+1) * perRange,
		}
		m.Value = values[i]
		result = append(result, m)
	}

	return result
}
