package usage_window

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/rate_snapshot"
	"time"
)

func New(
	start time.Time,
	end time.Time,
	rate []rate_snapshot.Snapshot,
	fable []fable_snapshot.Snapshot,
) *Window {
	return &Window{Start: start, End: end, Rate: rate, Fable: fable}
}
