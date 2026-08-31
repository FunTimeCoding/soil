package goatlassiand

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence"
	"github.com/funtimecoding/soil/pkg/atlassian/jira"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/option"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/web"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	webLib "github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Atlassian,
	s face.Instrument,
) {
	r := s.Reporter()
	j := jira.NewEnvironment()
	c := confluence.NewEnvironment()
	l := logger.New(context.Background())
	k := worker.New(j, c, constant.PollInterval, l, r)
	b := web.New(k)
	lifecycle.New(
		l,
		lifecycle.WithWorker(k),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					t := s.Recorder()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(j, c, r),
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
										webLib.RecordTelemetry(t, operation, e)

										return response, e
									}
								},
							},
						),
						m,
					)
					model_context.New(j, c, r, t, o.Version).Mount(m)
					b.Mount(m)
				},
			).WithMiddleware(b.Recovery(r)),
		),
	).RunUntilSignal()
}
