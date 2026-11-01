package goqueryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"time"
)

func Run(
	o *option.Query,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	v := service.New(o, l, r)
	defer v.Close()
	u := web.New(v, r)
	c := lifecycle.New(l)

	for _, w := range v.Workers() {
		c.WithWorker(w)
	}

	e := metric.New()
	c.WithWorker(worker.New(v, 10*time.Minute, l, e.Registry(), r)).WithServer(
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
				Mount(v, u, r, i.Recorder(), guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(u.Recovery(r)),
	).RunUntilSignal()
}
