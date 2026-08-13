package issue

import (
	"sort"
	"time"
)

func SortByAge(v []*Issue) []*Issue {
	result := append([]*Issue{}, v...)
	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return time.Time(result[i].Raw.Fields.Created).Before(
				time.Time(result[j].Raw.Fields.Created),
			)
		},
	)

	return result
}
