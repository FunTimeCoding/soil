package chunk

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"sort"
)

func scanBreakPoints(text string) []breakPoint {
	seen := map[int]breakPoint{}

	for _, p := range constant.BreakPatterns {
		for _, match := range p.Pattern.FindAllStringIndex(text, -1) {
			position := match[0]
			existing, found := seen[position]

			if !found || p.Score > existing.score {
				seen[position] = breakPoint{
					position: position,
					score:    p.Score,
				}
			}
		}
	}

	result := make([]breakPoint, 0, len(seen))

	for _, b := range seen {
		result = append(result, b)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].position < result[j].position
		},
	)

	return result
}
