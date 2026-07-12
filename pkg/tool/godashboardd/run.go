package godashboardd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/worker"
	"net/http"
)

func Run(
	o *option.Dashboard,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	c := store.New(relational.Open(l, o.PostgresLocator, o.LitePath))
	defer c.Close()
	v := service.New(
		o.Board,
		prometheus.New(
			o.Board.Connection.Prometheus.Host,
			board.Port(o.Board.Connection.Prometheus),
			o.Board.Connection.Prometheus.Secure,
			"",
			"",
			"",
		),
		usageClient(o.Board),
		argocdClient(o.Board),
		notifier.New(),
		l,
	)
	u := web.New(
		o.Board,
		v,
		c,
		authorizationClient(o),
	)
	lifecycle.New(
		l,
		lifecycle.WithWorker(
			worker.New(v, constant.RefreshInterval, l, r),
		),
		lifecycle.WithServer(
			server.New(
				o.Address,
				func(m *http.ServeMux) {
					u.Mount(m)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
