package gonetboxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Netbox,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(relational.Open(l, o.PostgresLocator, o.LitePath))
	defer s.Close()
	m := metric.New()
	c := netbox.NewEnvironment(
		netbox.WithRequestDuration(m.Upstream(constant.Identity)),
	)
	defer c.Close()
	b := web.New(c)
	lifecycle.New(l).WithServer(
		server.New(
			constant.Identity,
			o.MetricAddress,
			func(x *http.ServeMux) {
				x.Handle(webConstant.MetricsPath, m.Exporter())
			},
		),
	).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				Mount(
					c,
					service.New(c, s),
					b,
					r,
					i.Recorder(),
					guard.New(m, o.ServiceTokens),
				)
			},
		).WithMiddleware(b.Recovery(r)),
	).RunUntilSignal()
}
