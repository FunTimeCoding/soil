package bluetooth

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"time"
)

type Collector struct {
	store    *store.Store
	logger   *logger.Logger
	recovery *recovery.Recovery
	interval time.Duration
	stop     chan struct{}
	last     map[string]string
}
