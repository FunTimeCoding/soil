package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
)

type Watcher struct {
	service  *service.Service
	recovery *recovery.Recovery
	logger   *logger.Logger
	harbor   string
	changed  map[string]string
	notifier *fsnotify.Watcher
}
