package matched

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
	"slices"
)

func (m *Matched) Current(
	v *system.System,
	trackers []*tracker.Tracker,
	scores uint64,
) bool {
	return m.Updated.Equal(v.UpdatedAt) &&
		slices.Equal(m.Trackers, trackers) &&
		m.Scores == scores
}
