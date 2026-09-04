package goflightd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/bluetooth"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/stream"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/collector/wireless"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/janitor"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"time"
)

func Run(
	o *option.Flight,
	i face.Instrument,
) {
	r := i.Reporter()
	g := logger.New(context.Background())
	s := store.New(relational.Open(g, o.PostgresLocator, o.LitePath))
	defer s.Close()
	options := []lifecycle.Option{
		lifecycle.WithWorker(stream.New(s, g, r, o.Predicate)),
		lifecycle.WithWorker(bluetooth.New(s, g, r, 15*time.Second)),
		lifecycle.WithWorker(janitor.New(s, g, r, time.Hour, 7*24*time.Hour)),
	}
	w := wireless.New(s, g, r, 5*time.Second)

	if w.Probe() {
		options = append(options, lifecycle.WithWorker(w))
	} else {
		g.Structured(
			"wireless collector disabled",
			"reason",
			"wdutil requires passwordless sudo",
		)
	}

	options = append(
		options,
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
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
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	)
	lifecycle.New(g, options...).RunUntilSignal()
}
