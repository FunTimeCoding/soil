package matched

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/tracker"
)

func New(
	v *system.System,
	trackers []*tracker.Tracker,
	scores uint64,
) *Matched {
	return &Matched{Updated: v.UpdatedAt, Trackers: trackers, Scores: scores}
}
