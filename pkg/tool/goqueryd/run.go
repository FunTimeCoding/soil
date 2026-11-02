package goqueryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Query,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	e := metric.New()
	d := daemon.New(o, l, r, e.Registry())
	defer d.Close()
	c := lifecycle.New(l)

	for _, w := range d.Workers() {
		c.WithWorker(w)
	}

	c.WithServer(
		server.New(
			constant.Identity,
			o.MetricAddress,
			func(x *http.ServeMux) {
				x.Handle(webConstant.MetricsPath, e.Exporter())
			},
		),
	).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				d.Mount(r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
