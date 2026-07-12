package goalertlogd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/telemetry"
	generated "github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	alertWeb "github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"time"
)

func Run(
	o *option.Log,
	r face.Reporter,
) {
	g := logger.New(context.Background())
	m := metric.New(0, false, nil)
	s := store.New(o.DatabasePath)
	defer s.Close()
	w := worker.New(
		alertmanager.NewEnvironment(),
		s,
		g,
		r,
		1*time.Minute,
		30*24*time.Hour,
		m.Registry(),
	)
	u := alertWeb.New(s, w)
	lifecycle.New(
		g,
		lifecycle.WithWorker(m),
		lifecycle.WithWorker(w),
		lifecycle.WithServer(
			lifecycleServer.New(
				o.Address,
				func(m *http.ServeMux) {
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, w, r),
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
					model_context.New(s, w, r, t, o.Version).Mount(m)
					u.Mount(m)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
