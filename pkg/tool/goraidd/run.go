package goraidd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/raid_parser"
	raidParserConstant "github.com/funtimecoding/soil/pkg/raid_parser/constant"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	raidWeb "github.com/funtimecoding/soil/pkg/tool/goraidd/web"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Raid,
	i face.Instrument,
) {
	r := i.Reporter()
	l := logger.New(context.Background())
	s := store.New(
		relational.Open(l, o.PostgresLocator, o.LitePath),
		o.LogCachePath,
		o.ElitePath,
		l,
		r,
	)
	p := raid_parser.New(
		"localhost:8081",
		true,
		environment.Required(raidParserConstant.TokenEnvironment),
	)
	u := raidWeb.New(s, o.ElitePath, o.OutputPath, p, authorizationClient(o))
	lifecycle.New(
		l,
		lifecycle.WithWorker(s),
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					t := i.Recorder()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, o.OutputPath, r),
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
						m,
					)
					u.Mount(guard.New(m, o.ServiceTokens))
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
