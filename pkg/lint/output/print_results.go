package output

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"slices"
)

func PrintResults(
	entries []*concern.Concern,
	summary bool,
) bool {
	var hasBlocked bool
	seen := make(map[string]bool)
	sorted := slices.Clone(entries)
	slices.SortStableFunc(
		sorted,
		func(
			a *concern.Concern,
			b *concern.Concern,
		) int {
			return cmp.Or(
				cmp.Compare(a.Path, b.Path),
				cmp.Compare(a.Line, b.Line),
			)
		},
	)

	for _, c := range sorted {
		line := FormatConcern(c)

		if !c.Fixed && !c.Planned {
			hasBlocked = true
			console.Line(line)

			continue
		}

		if summary {
			if !seen[c.Path] {
				seen[c.Path] = true
				console.Line(c.Path)
			}

			continue
		}

		console.Line(line)
	}

	return hasBlocked
}
