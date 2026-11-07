package gocredentiald

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Credential,
	i face.Instrument,
) {
	r := i.Reporter()
	t := i.Recorder()
	l := logger.New(context.Background())
	d := daemon.New(o, l, r)
	lifecycle.New(l).WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, t, g)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).RunUntilSignal()
}
