package hold

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func New(
	name string,
	c face.Connection,
	l *logger.Logger,
	r face.Reporter,
) *Hold {
	x, cancel := context.WithCancel(context.Background())

	return &Hold{
		name:       name,
		connection: c,
		reacher:    reacher.New(r),
		recovery:   recovery.New(l, r),
		edge:       func(e *edge.Edge) { l.Plain("%s", e) },
		refusal:    func(e error) { r.CaptureException(e) },
		minimum:    constant.ProbeMinimum,
		maximum:    constant.ProbeMaximum,
		context:    x,
		cancel:     cancel,
		done:       make(chan struct{}),
	}
}
