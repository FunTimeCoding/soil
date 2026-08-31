package usage_window

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/rate_snapshot"
	"time"
)

type Window struct {
	Start time.Time
	End   time.Time
	Rate  []rate_snapshot.Snapshot
	Fable []fable_snapshot.Snapshot
}
