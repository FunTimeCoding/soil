package godashboardd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/worker"
	sharedWeb "github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func Run(
	o *option.Dashboard,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	c := store.New(o.PostgresLocator, o.LitePath)
	defer c.Close()
	target := o.Board.Connection.Prometheus
	v := service.New(
		o.Board,
		prometheus.New(
			target.Host,
			board.Port(target),
			target.Secure,
			"",
			"",
			"",
		),
		usageClient(o.Board),
		argocdClient(o.Board),
		notifier.New(),
		l,
	)
	lifecycle.New(
		l,
		lifecycle.WithWorker(
			worker.New(v, constant.RefreshInterval, l, r),
		),
		lifecycle.WithServer(
			server.New(
				sharedWeb.AddressPort(o.Port),
				func(m *http.ServeMux) {
					web.New(
						o.Board,
						v,
						c,
						authorizationClient(o),
					).Mount(m)
				},
			).WithMiddleware(sharedWeb.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
