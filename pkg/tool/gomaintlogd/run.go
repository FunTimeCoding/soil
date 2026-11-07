package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Run(
	o *option.Log,
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
				d.Mount(r, i.Recorder(), g)
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
