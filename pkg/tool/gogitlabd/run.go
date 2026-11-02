package gogitlabd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Gitlab,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	m := metric.New()
	d := daemon.New(o, l, r, m.Registry())
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
				x.Handle(webConstant.MetricsPath, m.Exporter())
			},
		),
	).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(x *http.ServeMux) {
				d.Mount(r, s.Recorder(), guard.New(x, o.ServiceTokens))
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
