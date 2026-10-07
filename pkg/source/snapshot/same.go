package snapshot

import (
	"github.com/funtimecoding/soil/pkg/source/types/snapshot_stamp"
	"maps"
)

func (s *Snapshot) Same(other *Snapshot) bool {
	return maps.EqualFunc(
		s.roots,
		other.roots,
		func(a map[string]snapshot_stamp.Stamp, b map[string]snapshot_stamp.Stamp) bool {
			return maps.EqualFunc(
				a,
				b,
				func(x snapshot_stamp.Stamp, y snapshot_stamp.Stamp) bool {
					return x.Size == y.Size && x.Modified.Equal(y.Modified)
				},
			)
		},
	)
}
