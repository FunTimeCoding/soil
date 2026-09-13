package goalertlogd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"time"
)

func Run(
	o *option.Log,
	i face.Instrument,
) {
	r := i.Reporter()
	g := logger.New(context.Background())
	m := metric.New()
	s := store.New(relational.Open(g, o.PostgresLocator, o.LitePath))
	defer s.Close()
	events := notifier.New()
	w := worker.New(
		alertmanager.NewEnvironment(),
		s,
		events,
		g,
		r,
		1*time.Minute,
		30*24*time.Hour,
		m.Registry(),
	)
	u := web.New(s, w, events)
	lifecycle.New(
		g,
		lifecycle.WithWorker(w),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.MetricAddress,
				func(x *http.ServeMux) {
					x.Handle(webConstant.MetricsPath, m.Exporter())
				},
			),
		),
		lifecycle.WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(m *http.ServeMux) {
					Mount(
						s,
						w,
						u,
						r,
						i.Recorder(),
						o.Version,
						guard.New(m, o.ServiceTokens),
					)
				},
			).WithMiddleware(u.Recovery(r)),
		),
	).RunUntilSignal()
}
