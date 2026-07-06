package gofix

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
)

func printSurvey(
	counts map[string]int,
	examples map[string][]string,
) {
	type entry struct {
		segment string
		count   int
	}
	var entries []entry

	for segment, count := range counts {
		entries = append(entries, entry{segment, count})
	}

	sort.Slice(
		entries,
		func(i, j int) bool {
			return entries[i].count > entries[j].count
		},
	)

	for _, e := range entries {
		fmt.Printf(
			"%4d  %-20s  %s\n",
			e.count,
			e.segment,
			join.CommaSpace(examples[e.segment]),
		)
	}
}
