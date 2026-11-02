package gotelemetryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Telemetry,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	d := daemon.New(o, l, r)
	defer d.Close()
	lifecycle.New(l).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(m *http.ServeMux) {
				d.Mount(r, guard.New(m, o.ServiceTokens))
			},
		).WithMiddleware(d.Recovery(r)),
	).RunUntilSignal()
}
