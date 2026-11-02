package hold

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"sync"
	"time"
)

type Hold struct {
	mutex      sync.Mutex
	serving    bool
	name       string
	connection face.Connection
	reacher    *reacher.Reacher
	recovery   *recovery.Recovery
	edge       func(*edge.Edge)
	refusal    func(error)
	minimum    time.Duration
	maximum    time.Duration
	context    context.Context
	cancel     context.CancelFunc
	started    sync.Once
	done       chan struct{}
}
