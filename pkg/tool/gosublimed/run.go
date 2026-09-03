package gosublimed

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/sublime"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/option"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Sublime,
	s face.Instrument,
) {
	r := s.Reporter()
	c := sublime.NewEnvironment()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					t := s.Recorder()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(c),
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
					model_context.New(c, r, t, o.Version).Mount(
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
