package goalpined

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Alpine,
	i face.Instrument,
) {
	r := i.Reporter()
	t := i.Recorder()
	l := logger.New(context.Background())
	d := daemon.New(o)
	lifecycle.New(l).WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, t, g)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).WithServer(
		server.New(constant.Identity, &o.Files, r, d.MountFiles),
	).RunUntilSignal()
}
