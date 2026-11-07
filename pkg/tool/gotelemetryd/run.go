package gotelemetryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Telemetry,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	d := daemon.New(o, l, r)
	defer d.Close()
	lifecycle.New(l).WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, g)
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
