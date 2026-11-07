package goagentd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Agent,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	d := daemon.New(o, l, r)
	defer d.Close()
	c := lifecycle.New(l)
	c.WithWorkers(d.Workers()...)
	c.WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, i.Recorder(), g)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)).WithProtected(),
	).RunUntilSignal()
}
