package goatlasd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/netbox"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Atlas,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	m := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(m)
	s := store.New(m)
	n := netbox.NewEnvironment()
	e := metric.New()
	u := web.New(s, authorizationClient(o))
	lifecycle.New(
		l,
		lifecycle.WithWorker(
			worker.New(
				s,
				collectors(n, l),
				lease.NewOptional(),
				n,
				o.Interval,
				o.Retention,
				l,
				e.Registry(),
				r,
			),
		),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.MetricAddress,
				func(x *http.ServeMux) {
					x.Handle(webConstant.MetricsPath, e.Exporter())
				},
			),
		),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(x *http.ServeMux) {
					Mount(s, u, r, i.Recorder(), guard.New(x, o.ServiceTokens))
				},
			).WithMiddleware(u.Recovery(r)).WithProtected(),
		),
	).RunUntilSignal()
}
