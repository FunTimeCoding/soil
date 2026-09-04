package gopnsensed

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/opnsense"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/server"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Opnsense,
	s face.Instrument,
) {
	r := s.Reporter()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					c := opnsense.NewEnvironment()
					t := s.Recorder()
					guard.New(m, o.ServiceTokens).TokenMount(
						webConstant.InterfacePath,
						generated.HandlerFromMux(
							generated.NewStrictHandler(
								server.New(c, r),
								[]generated.StrictMiddlewareFunc{
									func(
										f generated.StrictHandlerFunc,
										operation string,
									) generated.StrictHandlerFunc {
										return func(
											x context.Context,
											w http.ResponseWriter,
											r *http.Request,
											request any,
										) (any, error) {
											response, e := f(x, w, r, request)
											web.RecordTelemetry(t, operation, e)

											return response, e
										}
									},
								},
							),
							http.NewServeMux(),
						),
					)
					model_context.New(c, r, t, o.Version).Mount(
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
