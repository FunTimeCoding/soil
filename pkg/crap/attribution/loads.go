package attribution

import (
	"github.com/funtimecoding/soil/pkg/crap/types/load"
	"sort"
)

func (m *Matrix) Loads() []*load.Load {
	alone := map[string]int{}
	total := map[string]int{}

	for _, tests := range m.Covers {
		for _, t := range tests {
			total[t]++
		}

		if len(tests) == 1 {
			alone[tests[0]]++
		}
	}

	result := make([]*load.Load, 0, len(m.Tests))

	for _, t := range m.Tests {
		result = append(result, load.New(t, total[t], alone[t]))
	}

	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			if result[i].Alone != result[j].Alone {
				return result[i].Alone > result[j].Alone
			}

			if result[i].Total != result[j].Total {
				return result[i].Total > result[j].Total
			}

			return result[i].Test < result[j].Test
		},
	)

	return result
}
