package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/worker"
)

func clocks(
	v *service.Service,
	o *option.Option,
	l *logger.Logger,
	r face.Reporter,
) []face.Worker {
	result := []face.Worker{
		worker.NewOnce(v.Warm, l, r),
		worker.NewOnce(v.CatchUpSearch, l, r),
		worker.New(v.Reconcile, l, r, constant.ReconcileInterval),
		interim(v.RunTimeoutSweep, l, r, constant.TimeoutSweepInterval),
		interim(v.PollMemory, l, r, constant.MemoryPollInterval),
		interim(v.DeliverImpressions, l, r, constant.ImpressionInterval),
		interim(v.ReleaseCruise, l, r, constant.CruiseInterval),
	}

	if o.MonitorUsage {
		result = append(
			result,
			interim(v.PollUsage, l, r, constant.UsagePollInterval),
		)
	}

	if o.UtilizationFallback {
		result = append(
			result,
			interim(v.PollUtilization, l, r, constant.UtilizationPollInterval),
		)
	}

	return result
}
