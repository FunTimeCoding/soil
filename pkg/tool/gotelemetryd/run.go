package gotelemetryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Telemetry,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(relational.Open(l, o.PostgresLocator, o.LitePath))
	defer s.Close()
	u := web.New(s)
	lifecycle.New(
		l,
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						s,
						u,
						r,
						i.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
