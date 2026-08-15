package stream

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
)

type Collector struct {
	store     *store.Store
	logger    *logger.Logger
	recovery  *recovery.Recovery
	predicate string
	stop      chan struct{}
}
