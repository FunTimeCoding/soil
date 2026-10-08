package gofix

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maps"
	"slices"
	"sort"
)

func printSurvey(
	counts map[string]int,
	examples map[string][]string,
) {
	segments := slices.Collect(maps.Keys(counts))
	sort.Slice(
		segments,
		func(i, j int) bool {
			return counts[segments[i]] > counts[segments[j]]
		},
	)

	for _, s := range segments {
		console.Format(
			"%4d  %-20s  %s\n",
			counts[s],
			s,
			join.CommaSpace(examples[s]),
		)
	}
}
