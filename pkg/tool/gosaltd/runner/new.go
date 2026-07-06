package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/option"
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
