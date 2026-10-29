package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"sync"
)

type Runner struct {
	provision     *runner.Runner
	store         *store.Store
	saltConnector func() *salt.Client
	salt          *salt.Client
	built         *salt.Client
	mutex         sync.Mutex
	closed        bool
	connected     chan struct{}
	logger        *logger.Logger
	reporter      face.Reporter
	recovery      *recovery.Recovery
	reacher       *reacher.Reacher
}
