package goclauded

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Option,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	m := metric.New()
	d := daemon.New(o, l, r, m.Registry())
	defer d.Close()
	c := lifecycle.New(l)
	c.WithWorkers(d.Workers()...)
	c.WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, constant.ChannelHold, i.Recorder(), g)
			},
		).WithMiddleware(d.Recovery(r)).
			WithProfiling().
			WithDefaultCertificate(),
	).WithServer(
		server.New(constant.Identity, &o.Metric, r, m.Mount),
	).WithWorker(d.Stopper()).RunUntilSignal()
}
