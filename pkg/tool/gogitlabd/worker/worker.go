package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

type Worker struct {
	client   *gitlab.Client
	interval time.Duration
	gauge    *prometheus.GaugeVec
	recovery *recovery.Recovery
	notifier *notifier.Notifier
	stop     chan struct{}
	mutex    sync.RWMutex
	entries  []*board_entry.Entry
}
