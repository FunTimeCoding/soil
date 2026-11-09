package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	saltFace "github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
)

func New(
	o *option.Salt,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	registry face.ProcessRegistry,
	c saltFace.SaltClient,
	g *peer.Guard,
) *Runner {
	result := &Runner{
		store:           s,
		salt:            c,
		guard:           g,
		keySyncInterval: o.KeySyncInterval,
		logger:          l,
		reporter:        r,
		recovery:        recovery.New(l, r),
	}
	result.provision = runner.New(
		runner_option.Option{
			Repository:      o.Repository,
			CloneDirectory:  o.CloneDirectory,
			CopyModes:       o.CopyModes,
			ToolPath:        o.SaltDirectory,
			ApplyFunction:   result.apply,
			SetupFunction:   result.awaitMaster,
			CleanupFunction: s.Cleanup,
			Registry:        registry,
		},
		l,
		r,
	)

	return result
}
