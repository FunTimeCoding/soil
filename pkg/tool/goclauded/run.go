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
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/telemetry"
	"github.com/funtimecoding/soil/pkg/ticker"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_indexer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/sweep"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	memory "github.com/funtimecoding/soil/pkg/tool/gomemoryd/connect"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/connect"
	library "github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"time"
)

func Run(
	o *option.Option,
	r face.Reporter,
) {
	start := time.Now()
	elapsed := func() float64 { return time.Since(start).Seconds() }
	l := logger.New(context.Background())
	n := notifier.New()
	s := store.New(o.LitePath, time.Now)
	l.Structured("store_ready", "elapsed", elapsed())
	h := claude.New().Base()
	result := sweep.Run(h)
	l.Structured(
		"sweep_complete",
		"elapsed",
		elapsed(),
		"copied",
		result.Copied,
		"updated",
		result.Updated,
		"skipped",
		result.Skipped,
	)
	memoryClient := memory.Wait(l)
	l.Structured("gomemoryd_connected", "elapsed", elapsed())
	queryClient := connect.Wait(l)
	l.Structured("goqueryd_connected", "elapsed", elapsed())
	summaryIdx := session_indexer.New(
		queryClient,
		"summaries",
		"session-summary",
	)
	completionIdx := session_indexer.New(
		queryClient,
		"completions",
		"session-completion",
	)
	v := service.New(
		s,
		claude.New(),
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
	l.Structured("summaries_reconciled", "elapsed", elapsed())
	v.ReconcileCompletions()
	l.Structured("completions_reconciled", "elapsed", elapsed())
	v.PopulateCache()
	l.Structured("cache_populated", "elapsed", elapsed())
	v.BackfillSessions()
	l.Structured("sessions_backfilled", "elapsed", elapsed())
	v.CheckConsistency()
	l.Structured("consistency_checked", "elapsed", elapsed())
	rec := recovery.New(l, r)
	timeoutTicker := ticker.New(
		5*time.Minute,
		v.RunTimeoutSweep,
		rec,
	)
	memoryTicker := ticker.New(30*time.Second, v.PollMemory, rec)
	w := watcher.New(v, l, r, h)
	address := o.Address
	t := telemetry.NewEnvironment()
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
		model_context.New(
			v,
			r,
			l,
			t,
			o.Version,
		).Mount(m)
		web.New(v).Mount(m)
	}
	middleware := library.RecoveryMiddleware(r)
	srv := lifecycleServer.New(address, setup).
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
			lifecycle.WithWorker(
				ticker.New(time.Minute, v.PollUsage, rec),
			),
		)
	}

	lifecycle.New(l, options...).RunUntilSignal()
}
