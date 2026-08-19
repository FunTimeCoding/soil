package gocertificated

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/telemetry"
	generated "github.com/funtimecoding/soil/pkg/tool/gocertificated/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/server"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	certificateWeb "github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Certificate,
	r face.Reporter,
) {
	g := logger.New(context.Background())
	s := store.New(relational.Open(g, o.PostgresLocator, o.LitePath))
	defer s.Close()
	f := gitlab.NewEnvironment()
	project, e := f.ResolveProject(o.Project)
	errors.PanicOnError(e)
	v := service.New(
		s,
		publish.New(f, project, o.Branch).WithSecret(
			o.SecretAuthority,
			o.SecretPath,
		),
	)
	i := certificateWeb.New(s, v)
	lifecycle.New(
		g,
		lifecycle.WithServer(
			lifecycleServer.New(
				o.Address,
				func(m *http.ServeMux) {
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, v, r),
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
					model_context.New(s, v, r, t, o.Version).Mount(m)
					i.Mount(m)
				},
			).WithMiddleware(i.Recovery(r)),
		),
	).RunUntilSignal()
}
