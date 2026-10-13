package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

type Worker struct {
	service  *service.Service
	interval time.Duration
	gauge    *prometheus.GaugeVec
	reporter face.Reporter
	recovery *recovery.Recovery
	notifier *notifier.Notifier
	stop     chan struct{}
	mutex    sync.RWMutex
	entries  map[string][]*board_entry.Entry
	requests map[string][]*request_entry.Entry
}
