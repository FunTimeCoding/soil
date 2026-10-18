package peer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"time"
)

type Guard struct {
	name     string
	address  string
	probe    face.Reachable
	reacher  *reacher.Reacher
	logger   *logger.Logger
	reporter face.Reporter
	minimum  time.Duration
	maximum  time.Duration
	wake     chan struct{}
	context  context.Context
	cancel   context.CancelFunc
}
