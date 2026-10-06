package snapshot

import (
	"maps"
	"slices"
)

func (s *Snapshot) Diff(other *Snapshot) []string {
	changed := make(map[string]bool)

	for root, before := range s.roots {
		now := other.roots[root]

		for path, b := range before {
			n, okay := now[path]

			if !okay || n.size != b.size || !n.modified.Equal(b.modified) {
				changed[path] = true
			}
		}

		for path := range now {
			if _, okay := before[path]; !okay {
				changed[path] = true
			}
		}
	}

	return slices.Sorted(maps.Keys(changed))
}
