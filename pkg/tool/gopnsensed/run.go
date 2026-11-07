package gopnsensed

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Opnsense,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	d := daemon.New(o, l, r)
	defer d.Close()
	errors.PanicOnError(d.Wait(context.Background()))
	c := lifecycle.New(l)
	c.WithWorkers(d.Workers()...)
	c.WithServer(
		server.New(
			constant.Identity,
			&o.Web,
			r,
			func(g *guard.Mux) {
				d.Mount(r, s.Recorder(), g)
			},
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).RunUntilSignal()
}
