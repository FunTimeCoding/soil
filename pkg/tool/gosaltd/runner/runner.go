package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
)

type Runner struct {
	provision *runner.Runner
	store     *store.Store
	salt      *salt.Client
	connected chan struct{}
	logger    *logger.Logger
	reporter  face.Reporter
	recovery  *recovery.Recovery
	reacher   *reacher.Reacher
}
