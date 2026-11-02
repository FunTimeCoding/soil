package daemon

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

func New(
	o *option.Log,
	l *logger.Logger,
	r face.Reporter,
	m *metric.Metric,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	a := o.Alertmanager()
	n := notifier.New()
	result := &Daemon{
		database:     d,
		alertmanager: a,
		guard:        peer.New(constant.AlertmanagerPeer, a.Address(), a, l, r),
	}
	result.service = service.New(store.New(d), a, result.guard, o.Retention, n)
	result.web = web.New(result.service, n)
	result.clock = worker.New(
		result.service.Record,
		constant.PollInterval,
		l,
		m.Registry(),
		r,
	)

	return result
}
