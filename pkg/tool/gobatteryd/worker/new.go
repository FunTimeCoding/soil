package worker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/poll"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	read func() *poll.Poll,
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	battery := []string{constant.ServerLabel, constant.BatteryLabel}
	result := &Worker{
		read:     read,
		measures: map[string]*prometheus.GaugeVec{},
		status: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: constant.StatusMetric,
				Help: constant.StatusHelp,
			},
			append(battery, constant.FlagLabel),
		),
		information: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: constant.InformationMetric,
				Help: constant.InformationHelp,
			},
			append(
				battery,
				constant.ModelLabel,
				constant.ManufacturerLabel,
				constant.SerialLabel,
			),
		),
		attached: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: constant.AttachedMetric,
				Help: constant.AttachedHelp,
			},
			battery,
		),
		reachable: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: constant.ReachableMetric,
				Help: constant.ReachableHelp,
			},
			[]string{constant.ServerLabel},
		),
	}

	for _, m := range constant.Measures {
		result.measures[m.Variable] = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{Name: m.Name, Help: m.Help},
			battery,
		)
		y.MustRegister(result.measures[m.Variable])
	}

	y.MustRegister(
		result.status,
		result.information,
		result.attached,
		result.reachable,
	)
	result.clock = worker.NewEvery(result.tick, l, r, interval)

	return result
}
