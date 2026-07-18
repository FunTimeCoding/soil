package gonetboxd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/telemetry"
	generated "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Netbox,
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
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(o.Client, s, r),
							[]generated.StrictMiddlewareFunc{
								func(
									f generated.StrictHandlerFunc,
									operation string,
								) generated.StrictHandlerFunc {
									return func(
										x context.Context,
										w http.ResponseWriter,
										q *http.Request,
										request any,
									) (any, error) {
										response, e := f(x, w, q, request)
										web.RecordTelemetry(t, operation, e)

										return response, e
									}
								},
							},
						),
						m,
					)
					model_context.New(
						o.Client,
						s,
						r,
						t,
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
