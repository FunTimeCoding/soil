package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
	"sync"
	"time"
)

type Worker struct {
	refresh   func() []*sample.Sample
	interval  time.Duration
	collector *collector.Collector
	log       *logger.Logger
	reporter  face.Reporter
	recovery  *recovery.Recovery
	stop      chan struct{}
	done      chan struct{}
	started   sync.Once
	stopped   sync.Once
}
