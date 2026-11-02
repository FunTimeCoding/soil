package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/time/time_of_day"
)

func NewDaily(
	tick func() []error,
	l *logger.Logger,
	r face.Reporter,
	at *time_of_day.Time,
) *Worker {
	result := New(func(context.Context) []error { return tick() }, l, r, 0)
	result.at = at

	return result
}
