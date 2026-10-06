package snapshot

import "slices"

func (s *Snapshot) Changed(roots ...string) []string {
	var result []string

	for _, r := range roots {
		before := s.roots[r]
		now := walk(r)

		for path, n := range now {
			b, okay := before[path]

			if !okay || b.size != n.size || !b.modified.Equal(n.modified) {
				result = append(result, path)
			}
		}

		for path := range before {
			if _, okay := now[path]; !okay {
				result = append(result, path)
			}
		}
	}

	return slices.Sorted(slices.Values(result))
}
