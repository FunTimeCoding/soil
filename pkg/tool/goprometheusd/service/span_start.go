package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"time"
)

func spanStart(
	now time.Time,
	span time.Duration,
) time.Time {
	result, _ := window.Whole().Bound(now, span).Start()

	return result
}
