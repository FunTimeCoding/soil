package gotelemetryd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/telemetry"
	generated "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	telemetryWeb "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Telemetry,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(relational.Open(l, o.PostgresLocator, o.LitePath))
	defer s.Close()
	lifecycle.New(
		l,
		lifecycle.WithServer(
			lifecycleServer.New(
				o.Address,
				func(m *http.ServeMux) {
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, r),
							nil,
						),
						m,
					)
					model_context.New(
						service.New(s),
						r,
						telemetry.NewEnvironment(),
						o.Version,
					).Mount(m)
					telemetryWeb.New(s).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
