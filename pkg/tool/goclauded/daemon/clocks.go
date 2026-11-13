package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/cruise"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/document"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/funtimecoding/soil/pkg/worker"
)

func clocks(
	v *roster.Service,
	t *transcript.Service,
	c *cruise.Service,
	y *memory.Service,
	u *usage.Service,
	x *search.Service,
	d *document.Service,
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
) []face.Worker {
	result := []face.Worker{
		worker.NewOnce(t.Warm, l, r),
		worker.NewOnce(x.CatchUpSearch, l, r),
		worker.New(d.Reconcile, l, r, constant.ReconcileInterval),
		worker.NewEvery(
			v.RunTimeoutSweep,
			l,
			r,
			constant.TimeoutSweepInterval,
		).WithDelayedStart(),
		worker.NewEvery(y.PollMemory, l, r, constant.MemoryPollInterval),
		worker.NewEvery(y.DeliverImpressions, l, r, constant.ImpressionInterval),
		worker.NewEvery(c.Release, l, r, constant.CruiseInterval),
		worker.NewEvery(
			t.RefreshReports,
			l,
			r,
			constant.ReportInterval,
		).WithDelayedStart(),
	}

	if o.MonitorUsage {
		result = append(
			result,
			worker.NewEvery(u.PollUsage, l, r, constant.UsagePollInterval),
		)
	}

	if o.UtilizationFallback {
		result = append(
			result,
			worker.New(
				u.PollUtilization,
				l,
				r,
				constant.UtilizationPollInterval,
			),
		)
	}

	return result
}
