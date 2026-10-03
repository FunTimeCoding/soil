package gofix

import (
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"sort"
)

func formatFile(
	name string,
	source []byte,
) (*formattedFile, error) {
	_, originalLines := significantTokens(source)
	result := &formattedFile{Source: source, Converged: true}
	seen := make(map[int]bool)

	for _, collapse := range []bool{true, false} {
		converged := false

		for range constant.MaxFormatPasses {
			changes, next, e := formatPass(name, result.Source, collapse)

			if e != nil {
				return nil, e
			}

			if len(changes) == 0 {
				converged = true

				break
			}

			offsets, _ := significantTokens(result.Source)

			for _, c := range changes {
				ordinal := sort.SearchInts(offsets, c.Offset)

				if ordinal >= len(offsets) || offsets[ordinal] != c.Offset ||
					seen[ordinal] {
					continue
				}

				seen[ordinal] = true
				c.Line = originalLines[ordinal]
				result.Changes = append(result.Changes, c)
			}

			result.Source = next
		}

		if !converged {
			result.Converged = false
		}
	}

	return result, nil
}
