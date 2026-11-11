package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/metric"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/cruise"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/document"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/startup"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
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
	n := o.Notifier()
	migrate.AutoMigrate(d.database)
	s := store.New(d.database, o.Clock)
	x := search_index.New(d.searchDatabase)
	m := metric.New(y)
	d.label = label.New(s, n)
	d.message = message.New(s, n, d.stop, o.Clock, m, l)
	d.cruise = cruise.New(s, n, o.Clock)

	if o.ChannelAlways {
		d.message.WithChannelAlways()
		d.cruise.WithDefaultGrace(constant.DefaultCruiseGrace)
	}

	d.memory = memory.New(s, d.memoryPeer, d.memoryGuard, d.message, o.Clock, r)
	d.usage = usage.New(s, n, d.browser, o.Clock, l, r)
	d.search = search.New(s, x, o.Harbor)
	d.document = document.New(s, d.query, d.queryGuard, m)
	d.event = event.New(s, d.message, d.memory, d.document, n, o.Clock, r)
	c := o.Claude()
	cache := session_cache.New()
	d.transcript = transcript.New(
		s,
		c,
		cache,
		d.memory,
		x,
		o.Harbor,
		o.Sources,
		o.SessionExportDirectory,
		o.Clock,
		l,
		r,
	)

	if d.utilization != nil {
		d.usage.WithUtilization(d.utilization, d.utilizationGuard, o.Credential)
	}

	d.session = session.New(
		s,
		d.label,
		d.message,
		d.cruise,
		d.transcript,
		c,
		cache,
		x,
		n,
		o.Clock,
		o.Harbor,
		o.Sources,
	)
	d.roster = roster.New(s, d.message, d.session, cache, n, o.Clock, l)
	d.startup = startup.New(d.roster, d.transcript, d.search)
	d.web = web.New(
		d.roster,
		d.session,
		d.label,
		d.message,
		d.cruise,
		d.event,
		d.transcript,
		d.usage,
		d.search,
		n,
		r,
	)
	d.gauge = gauge.New(d.roster.Findings, constant.FindingInterval, l, y, r)
	d.watcher = watcher.New(
		d.startup.Rescan,
		d.transcript.SyncChanged,
		o.Sources,
		o.Harbor,
		l,
		r,
	)
	d.clocks = clocks(
		d.roster,
		d.transcript,
		d.cruise,
		d.memory,
		d.usage,
		d.search,
		d.document,
		o,
		l,
		r,
	)
}
