package goraidparsed

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Parser,
	s face.Instrument,
) {
	r := s.Reporter()
	l := logger.New(context.Background())
	d := daemon.New(o, l)
	lifecycle.New(l).WithServer(
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
