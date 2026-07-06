package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
)

func New(
	o *option.Salt,
	connector func() *salt.Client,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	registry face.ProcessRegistry,
) *Runner {
	result := &Runner{
		store:         s,
		saltConnector: connector,
		connected:     make(chan struct{}),
		logger:        l,
		recovery:      recovery.New(l, r),
	}
	result.provision = runner.New(
		runner.Configuration{
			Repository:      o.Repository,
			ClonePath:       o.ClonePath,
			ToolPath:        o.SaltPath,
			ApplyFunction:   result.apply,
			SetupFunction:   result.connectLoop,
			CleanupFunction: s.Cleanup,
			Registry:        registry,
		},
		l,
		r,
	)

	return result
}
