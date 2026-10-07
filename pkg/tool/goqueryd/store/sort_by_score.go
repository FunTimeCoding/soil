package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"sort"
)

func sortByScore(v []result.Ranked) {
	sort.Slice(
		v,
		func(i, j int) bool {
			return v[i].Score > v[j].Score
		},
	)
}
