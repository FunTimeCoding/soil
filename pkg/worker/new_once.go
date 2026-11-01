package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func NewOnce(
	tick func(context.Context) []error,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	result := New(tick, l, r, 0)
	result.once = true

	return result
}
