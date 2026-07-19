package output

import (
	"fmt"
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

		if !c.Fixed {
			hasBlocked = true
			fmt.Println(line)

			continue
		}

		if summary {
			if !seen[c.Path] {
				seen[c.Path] = true
				fmt.Println(c.Path)
			}

			continue
		}

		fmt.Println(line)
	}

	return hasBlocked
}
