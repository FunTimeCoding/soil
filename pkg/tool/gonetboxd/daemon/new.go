package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/worker"
)

func New(
	o *option.Netbox,
	l *logger.Logger,
	r face.Reporter,
	m *metric.Metric,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	c := o.Source(m.Upstream(constant.Identity))
	result := &Daemon{
		database: d,
		client:   c,
		service:  service.New(store.New(d), c),
	}
	result.web = web.New(result.service)

	if o.PowerIntent {
		result.clock = worker.New(
			result.service.Intents,
			constant.IntentInterval,
			l,
			m.Registry(),
			r,
		)
	}

	return result
}
