package gocertificated

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Certificate,
	u face.Instrument,
) {
	r := u.Reporter()
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
				d.Mount(r, u.Recorder(), g)
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
