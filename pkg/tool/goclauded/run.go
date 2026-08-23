package goclauded

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/ticker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/transcript_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	memory "github.com/funtimecoding/soil/pkg/tool/gomemoryd/connect"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/connect"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/indexer"
	library "github.com/funtimecoding/soil/pkg/web"
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
	n := notifier.New()
	s := store.New(lite.New(l, o.LitePath), time.Now)
	h := claude.New().Base()
	sweep.Run(h)
	memoryClient := memory.Wait(l)
	queryClient := connect.Wait(l)
	summaryIdx := indexer.New(queryClient, constant.SummarySourceType)
	completionIdx := indexer.New(queryClient, constant.CompletionSourceType)
	v := service.New(
		s,
		transcript_cache.New(claude.New()),
		memoryClient,
		summaryIdx,
		completionIdx,
		n,
		r,
		h,
		time.Now,
		l,
	)
	v.ClearBindings()
	v.ReconcileSummaries()
	v.ReconcileCompletions()
	v.PopulateCache()
	v.BackfillSessions()
	v.CheckConsistency()
	l.Structured("started", "elapsed", time.Since(start).Seconds())
	go warm(v)
	rec := recovery.New(l, r)
	timeoutTicker := ticker.New(5*time.Minute, v.RunTimeoutSweep, rec)
	memoryTicker := ticker.New(30*time.Second, v.PollMemory, rec)
	w := watcher.New(v, l, r, h)
	address := o.Address
	t := i.Recorder()
	u := web.New(v)
	setup := func(m *http.ServeMux) {
		generated.HandlerFromMux(
			generated.NewStrictHandler(
				server.New(v, l, r, h, o.SessionExportPath),
				[]generated.StrictMiddlewareFunc{
					func(
						f generated.StrictHandlerFunc,
						operation string,
					) generated.StrictHandlerFunc {
						return func(
							x context.Context,
							w http.ResponseWriter,
							q *http.Request,
							request any,
						) (any, error) {
							response, e := f(x, w, q, request)
							library.RecordTelemetry(t, operation, e)

							return response, e
						}
					},
				},
			),
			m,
		)
		model_context.New(v, r, l, t, o.Version).Mount(m)
		u.Mount(m)
	}
	middleware := u.Recovery(r)
	srv := lifecycleServer.New(constant.Identity, address, setup).
		WithMiddleware(middleware).
		WithProfiling().
		WithDefaultCertificate()
	options := []lifecycle.Option{
		lifecycle.WithWorker(w),
		lifecycle.WithServer(srv),
		lifecycle.WithWorker(timeoutTicker),
		lifecycle.WithWorker(memoryTicker),
	}

	if environment.Exists(constant.MonitorUsageEnvironment) {
		options = append(
			options,
			lifecycle.WithWorker(ticker.New(time.Minute, v.PollUsage, rec)),
		)
	}

	lifecycle.New(l, options...).RunUntilSignal()
}
