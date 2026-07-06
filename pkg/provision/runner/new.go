package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(
	c Configuration,
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
		registry:        c.Registry,
		logger:          l,
		recovery:        recovery.New(l, r),
		trigger:         make(chan TriggerRequest, 1),
		sync:            make(chan SyncRequest, 1),
		stop:            make(chan struct{}),
	}
}
