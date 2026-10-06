package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"sort"
)

func sortByScore(v []search.Ranked) {
	sort.Slice(
		v,
		func(i, j int) bool {
			return v[i].Score > v[j].Score
		},
	)
}
