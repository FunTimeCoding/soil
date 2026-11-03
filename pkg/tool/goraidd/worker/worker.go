package worker

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"sync"
)

type Worker struct {
	handle   func(context.Context, fsnotify.Event) []error
	paths    []string
	prepare  []func(context.Context) []error
	logger   *logger.Logger
	reporter face.Reporter
	recovery *recovery.Recovery
	context  context.Context
	cancel   context.CancelFunc
	running  sync.WaitGroup
}
