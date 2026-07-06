package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
)

type Runner struct {
	provision     *runner.Runner
	store         *store.Store
	saltConnector func() *salt.Client
	salt          *salt.Client
	connected     chan struct{}
	logger        *logger.Logger
	recovery      *recovery.Recovery
}
