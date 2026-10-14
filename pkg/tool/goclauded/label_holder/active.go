package label_holder

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func Active(
	t *target.Target,
	now time.Time,
	idleThreshold time.Duration,
) bool {
	if t.LastSeen.IsZero() {
		return true
	}

	return now.Sub(t.LastSeen) <= idleThreshold
}
