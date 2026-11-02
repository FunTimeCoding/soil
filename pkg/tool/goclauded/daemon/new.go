package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"time"
)

func New(
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	start := time.Now()
	m := o.Memory()
	q := o.Query()
	result := &Daemon{
		database: lite.New(l, o.LitePath),
		searchDatabase: connection.New(
			l,
			filepath.Join(filepath.Dir(o.LitePath), constant.SearchIndexFile),
		),
		memory:      m,
		memoryGuard: peer.New(constant.MemoryPeer, m.Address(), m, l, r),
		query:       q,
		queryGuard:  peer.New(constant.QueryPeer, q.Address(), q, l, r),
		stop:        stopper.New(),
		logger:      l,
	}

	if o.MonitorUsage {
		result.browser = o.Browser()
	}

	result.build(o, l, r, y)
	result.service.Prepare()
	l.Structured("started", "elapsed", time.Since(start).Seconds())

	return result
}
