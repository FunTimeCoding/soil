package goalertlogd

import (
	"context"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func Run(
	o *option.Log,
	i face.Instrument,
) {
	r := i.Reporter()
	g := logger.New(context.Background())
	m := metric.New()
	v := service.New(o, g, r)
	defer v.Close()
	events := notifier.New()
	w := worker.New(v, events, g, r, constant.PollInterval, m.Registry())
	u := web.New(v, w, events)
	c := lifecycle.New(g)

	for _, x := range v.Workers() {
		c.WithWorker(x)
	}

	c.WithWorker(w).WithServer(
		server.New(
			constant.Identity,
			o.MetricAddress,
			func(x *http.ServeMux) {
				x.Handle(webConstant.MetricsPath, m.Exporter())
			},
		),
	).WithServer(
		server.New(
			constant.Identity,
			o.Address,
			func(x *http.ServeMux) {
				Mount(v, w, u, r, i.Recorder(), guard.New(x, o.ServiceTokens))
			},
		).WithMiddleware(u.Recovery(r)),
	).RunUntilSignal()
}
