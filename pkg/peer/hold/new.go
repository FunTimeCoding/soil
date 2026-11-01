package hold

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
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
		logger:     l,
		reporter:   r,
		recovery:   recovery.New(l, r),
		minimum:    constant.ProbeMinimum,
		maximum:    constant.ProbeMaximum,
		context:    x,
		cancel:     cancel,
	}
}
