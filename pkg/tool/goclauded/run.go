package goclauded

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/ticker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/gauge"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/worker"
	"net/http"
	"time"
)

func Run(
	o *option.Option,
	i face.Instrument,
) {
	r := i.Reporter()
	start := time.Now()
	l := logger.New(context.Background())
	v := service.New(o, l, r)
	defer v.Close()
	v.Sweep()
	v.ClearBindings()
	v.PopulateCache()
	v.BackfillSessions()
	v.CheckConsistency()
	l.Structured("started", "elapsed", time.Since(start).Seconds())
	rec := recovery.New(l, r)
	m := metric.New()
	findingGauge := gauge.New(v, l, m.Registry())
	findingGauge.Poll()
	t := i.Recorder()
	u := web.New(v, r)
	c := lifecycle.New(l)

	for _, w := range v.Workers() {
		c.WithWorker(w)
	}

	c.WithWorker(worker.NewOnce(v.Warm, l, r)).
		WithWorker(worker.NewOnce(v.CatchUpSearch, l, r)).
		WithWorker(worker.New(v.Reconcile, l, r, constant.ReconcileInterval)).
		WithWorker(watcher.New(v, l, r, o.Harbor)).
		WithServer(
			server.New(
				constant.Identity,
				o.Address,
				func(x *http.ServeMux) {
					Mount(
						v,
						u,
						l,
						r,
						constant.ChannelHold,
						t,
						guard.New(x, o.ServiceTokens),
					)
				},
			).WithMiddleware(u.Recovery(r)).
				WithProfiling().
				WithDefaultCertificate(),
		).
		WithServer(
			server.New(
				constant.Identity,
				o.MetricAddress,
				func(x *http.ServeMux) {
					x.Handle(webConstant.MetricsPath, m.Exporter())
				},
			),
		).
		WithWorker(ticker.New(5*time.Minute, v.RunTimeoutSweep, rec)).
		WithWorker(ticker.New(30*time.Second, v.PollMemory, rec)).
		WithWorker(ticker.New(30*time.Second, v.DeliverImpressions, rec)).
		WithWorker(ticker.New(5*time.Minute, findingGauge.Poll, rec)).
		WithWorker(ticker.New(constant.CruiseInterval, v.ReleaseCruise, rec))

	if environment.Exists(constant.MonitorUsageEnvironment) &&
		!environment.Exists(constant.UsageEndpointEnvironment) {
		c.WithWorker(ticker.New(time.Minute, v.PollUsage, rec))
	}

	if v.UtilizationFallback() {
		c.WithWorker(ticker.New(5*time.Minute, v.PollUtilization, rec))
	}

	c.WithWorker(v.Stopper()).RunUntilSignal()
}
