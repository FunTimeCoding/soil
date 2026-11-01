package gonetboxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/worker"
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
	m := metric.New()
	v := service.New(o, l, m.Upstream(constant.Identity))
	defer v.Close()
	b := web.New(v)
	c := lifecycle.New(l)

	if o.PowerIntent {
		c.WithWorker(worker.New(v, constant.IntentInterval, l, m.Registry(), r))
	}

	c.WithServer(
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
			func(x *http.ServeMux) {
				Mount(v, b, r, i.Recorder(), guard.New(x, o.ServiceTokens))
			},
		).WithMiddleware(b.Recovery(r)),
	).RunUntilSignal()
}
