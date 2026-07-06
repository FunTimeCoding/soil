package parse_elite

import (
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"sort"
)

func sortByDamage(
	players map[string]*elite_parser.AggregatedPlayer,
) []*elite_parser.AggregatedPlayer {
	var result []*elite_parser.AggregatedPlayer

	for _, p := range players {
		result = append(result, p)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Offensive.Damage > result[j].Offensive.Damage
		},
	)

	return result
}
