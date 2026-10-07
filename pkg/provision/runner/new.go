package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

func New(
	c runner_option.Option,
	l *logger.Logger,
	r face.Reporter,
) *Runner {
	return &Runner{
		repository:      c.Repository,
		clonePath:       c.ClonePath,
		toolPath:        c.ToolPath,
		applyFunction:   c.ApplyFunction,
		initFunction:    c.InitFunction,
		setupFunction:   c.SetupFunction,
		cleanupFunction: c.CleanupFunction,
		changeFunction:  c.ChangeFunction,
		downstream:      c.Downstream,
		registry:        c.Registry,
		logger:          l,
		reporter:        r,
		recovery:        recovery.New(l, r),
		trigger:         make(chan trigger.Request, 1),
		sync:            make(chan update.Request, 1),
		stop:            make(chan struct{}),
	}
}
