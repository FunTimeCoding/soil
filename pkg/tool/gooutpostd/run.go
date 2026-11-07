package gooutpostd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Outpost,
	i face.Instrument,
) {
	r := i.Reporter()
	t := i.Recorder()
	d := daemon.New()
	lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, t, g)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)).WithProtected(),
	).RunUntilSignal()
}
