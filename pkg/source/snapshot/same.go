package snapshot

import "maps"

func (s *Snapshot) Same(other *Snapshot) bool {
	return maps.EqualFunc(
		s.roots,
		other.roots,
		func(a map[string]stamp, b map[string]stamp) bool {
			return maps.EqualFunc(
				a,
				b,
				func(x stamp, y stamp) bool {
					return x.size == y.size && x.modified.Equal(y.modified)
				},
			)
		},
	)
}
