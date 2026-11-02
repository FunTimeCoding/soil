package watcher

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"sync"
)

type Watcher struct {
	context     context.Context
	cancel      context.CancelFunc
	rescan      func(context.Context) []error
	sync        func(string, string)
	directories func() []string
	recovery    *recovery.Recovery
	reporter    face.Reporter
	harbor      string
	changed     map[string]string
	notifier    *fsnotify.Watcher
	started     sync.Once
	done        chan struct{}
}
