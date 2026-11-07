package goprocessd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Option,
	i face.Instrument,
) {
	r := i.Reporter()
	t := i.Recorder()
	d := daemon.New(o)
	defer d.Close()
	l := lifecycle.New(logger.New(context.Background()))
	l.WithWorkers(d.Workers()...)
	l.WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, t, g)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).WithServer(
		server.New(
			constant.Identity,
			&o.Local,
			r,
			func(g *guard.Mux) { d.MountLocal(r, t, g) },
		).WithListener(socket.Listen(socket.Path(o.ProcfilePath))).
			WithMiddleware(web.RecoveryMiddleware(r)),
	)
	l.Run()
	WaitForStop(d.Done())
	l.Close()
}
