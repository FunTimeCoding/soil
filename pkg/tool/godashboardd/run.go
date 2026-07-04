package godashboardd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/event/notifier"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/option"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/web"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/worker"
	sharedWeb "github.com/funtimecoding/go-library/pkg/web"
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
