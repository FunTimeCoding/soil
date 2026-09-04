package gopostgresd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Postgres,
	i face.Instrument,
) {
	r := i.Reporter()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					s := service.New(o.Inventory)
					t := i.Recorder()
					guard.New(m, o.ServiceTokens).TokenMount(
						webConstant.InterfacePath,
						generated.HandlerFromMux(
							generated.NewStrictHandler(
								server.New(s, r),
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
					model_context.New(s, r, t, o.Version).Mount(
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
