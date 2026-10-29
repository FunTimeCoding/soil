package goatlasd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
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
	e := metric.New()
	v := service.New(o, l, r)
	defer v.Close()
	u := web.New(v)
	c := lifecycle.New(l)

	for _, w := range v.Workers() {
		c.WithWorker(w)
	}

	c.WithWorker(worker.New(v, o.Interval, l, e.Registry(), r)).WithServer(
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
			func(x *http.ServeMux) {
				Mount(v, u, r, i.Recorder(), guard.New(x, o.ServiceTokens))
			},
		).WithMiddleware(u.Recovery(r)).WithProtected(),
	).RunUntilSignal()
}
