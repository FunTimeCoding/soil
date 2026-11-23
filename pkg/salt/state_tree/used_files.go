package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
	"strings"
)

func (t *Tree) usedFiles(live map[string]bool) map[string]bool {
	result := map[string]bool{constant.TopFile: true}
	references := slices.Clone(t.cited)
	patterns := slices.Clone(t.citedPattern)

	for s := range live {
		result[t.states[s]] = true
		references = append(references, t.references[s]...)
		patterns = append(patterns, t.filePattern[s]...)
	}

	for f := range t.files {
		if strings.HasPrefix(f, constant.PrivatePrefix) {
			result[f] = true

			continue
		}

		for _, r := range references {
			if f == r || strings.HasPrefix(f, join.Empty(r, "/")) {
				result[f] = true

				break
			}
		}

		for _, p := range patterns {
			if p.MatchString(f) {
				result[f] = true

				break
			}
		}
	}

	return result
}
