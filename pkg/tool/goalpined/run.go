package goalpined

import (
	"context"
	alpine "github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/alpine/package_server"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goalpined/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/server"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Alpine,
	i face.Instrument,
) {
	r := i.Reporter()
	t := i.Recorder()
	l := logger.New(context.Background())
	s := package_server.NewEnvironment()
	c := model_context.New(r, t, o.Version)
	lifecycle.New(
		l,
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					s.Mount(m)
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(r),
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
					c.Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				constant.FileAddress,
				func(m *http.ServeMux) {
					m.Handle(
						strings.Slash,
						http.FileServer(http.Dir(alpine.PackageRoot)),
					)
				},
			),
		),
	).RunUntilSignal()
}
