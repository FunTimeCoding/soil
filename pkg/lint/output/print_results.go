package output

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint/concern"
)

func PrintResults(
	entries []*concern.Concern,
	summary bool,
) bool {
	var hasBlocked bool
	seen := make(map[string]bool)

	for _, c := range entries {
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
