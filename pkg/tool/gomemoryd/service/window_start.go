package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"time"
)

func windowStart(
	now time.Time,
	span time.Duration,
) string {
	start, _ := window.Whole().Bound(now, span).Start()

	return start.UTC().Format(time.RFC3339)
}
