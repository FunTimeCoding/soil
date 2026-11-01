package hold

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"time"
)

type Hold struct {
	name       string
	connection face.Connection
	reacher    *reacher.Reacher
	logger     *logger.Logger
	reporter   face.Reporter
	recovery   *recovery.Recovery
	minimum    time.Duration
	maximum    time.Duration
	context    context.Context
	cancel     context.CancelFunc
}
