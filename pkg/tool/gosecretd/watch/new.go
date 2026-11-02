package watch

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/openbao/status"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	s func() (*status.Status, error),
	y *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
) *Watch {
	result := &Watch{
		status: s,
		sealed: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: constant.SealedGauge,
				Help: constant.SealedHelp,
			},
		),
		reachable: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: constant.ReachableGauge,
				Help: constant.ReachableHelp,
			},
		),
	}
	y.MustRegister(result.sealed, result.reachable)
	result.clock = worker.NewEvery(result.tick, l, r, constant.WatchInterval)

	return result
}
