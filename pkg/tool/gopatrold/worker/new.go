package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/types/inspection"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	round func(context.Context) ([]*inspection.Inspection, []error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	result := &Worker{round: round}
	host := []string{constant.HostLabel}
	array := []string{constant.HostLabel, constant.ArrayLabel}
	drive := []string{
		constant.HostLabel,
		constant.DriveLabel,
		constant.ModelLabel,
	}
	result.reached = result.gauge(
		y,
		constant.ReachedMetric,
		constant.ReachedHelp,
		host,
	)
	result.required = result.gauge(
		y,
		constant.RequiredMetric,
		constant.RequiredHelp,
		array,
	)
	result.active = result.gauge(
		y,
		constant.ActiveMetric,
		constant.ActiveHelp,
		array,
	)
	result.failed = result.gauge(
		y,
		constant.FailedMetric,
		constant.FailedHelp,
		array,
	)
	result.rebuilding = result.gauge(
		y,
		constant.RebuildingMetric,
		constant.RebuildingHelp,
		array,
	)
	result.passed = result.gauge(
		y,
		constant.PassedMetric,
		constant.PassedHelp,
		drive,
	)
	result.temperature = result.gauge(
		y,
		constant.TemperatureMetric,
		constant.TemperatureHelp,
		drive,
	)
	result.warning = result.gauge(
		y,
		constant.WarningMetric,
		constant.WarningHelp,
		drive,
	)
	result.spare = result.gauge(
		y,
		constant.SpareMetric,
		constant.SpareHelp,
		drive,
	)
	result.used = result.gauge(y, constant.UsedMetric, constant.UsedHelp, drive)
	result.media = result.gauge(
		y,
		constant.MediaMetric,
		constant.MediaHelp,
		drive,
	)
	result.clock = worker.New(result.tick, l, r, interval)

	return result
}
