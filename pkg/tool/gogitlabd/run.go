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
	"github.com/funtimecoding/soil/pkg/web/guard"
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
	c.WithWorkers(d.Workers()...)
	c.WithServer(
		server.New(constant.Identity, &o.Metric, r, m.Mount),
	).WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, s.Recorder(), g)
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
