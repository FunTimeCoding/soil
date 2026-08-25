package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"time"
)

type Worker struct {
	service   face.Service
	interval  time.Duration
	collector *collector.Collector
	log       *logger.Logger
	recovery  *recovery.Recovery
	stop      chan struct{}
}
