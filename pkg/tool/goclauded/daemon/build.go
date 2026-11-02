package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	"github.com/prometheus/client_golang/prometheus"
)

func (d *Daemon) build(
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) {
	v := service.New(
		store.New(d.database, o.Clock),
		search_index.New(d.searchDatabase),
		o.Claude(),
		d.memory,
		d.memoryGuard,
		d.query,
		d.queryGuard,
		o.Notifier(),
		d.browser,
		d.stop,
		o.Clock,
		o.SessionExportPath,
		o.Harbor,
		o.Sources,
		o.UtilizationFallback,
		l,
		r,
	)
	d.service = v
	d.web = web.New(v, r)
	d.gauge = gauge.New(v.Findings, constant.FindingInterval, l, y, r)
	d.watcher = watcher.New(v.Rescan, v.SyncChanged, o.Sources, o.Harbor, l, r)
	d.clocks = clocks(v, o, l, r)
}
