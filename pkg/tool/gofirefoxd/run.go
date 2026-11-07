package gofirefoxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Firefox,
	s face.Instrument,
) {
	r := s.Reporter()
	d := daemon.New(o)
	defer d.Close()
	lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(
			constant.Identity,
			&o.Bridge,
			r,
			d.MountBridge,
		).WithMiddleware(web.RecoveryMiddleware(r)),
	).WithServer(
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
