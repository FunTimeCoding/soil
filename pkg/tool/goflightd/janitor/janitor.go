package janitor

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"time"
)

type Janitor struct {
	store     *store.Store
	logger    *logger.Logger
	recovery  *recovery.Recovery
	interval  time.Duration
	retention time.Duration
	stop      chan struct{}
}
