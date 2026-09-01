package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"sync"
	"time"
)

type Worker struct {
	service   face.Service
	interval  time.Duration
	collector *collector.Collector
	log       *logger.Logger
	recovery  *recovery.Recovery
	notifier  *notifier.Notifier
	stop      chan struct{}
	mutex     sync.RWMutex
	floor     *floor.Floor
}
