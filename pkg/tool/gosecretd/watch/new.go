package watch

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/ticker"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	secret "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	s secret.SecretSource,
	y *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
) *Watch {
	result := &Watch{
		source: s,
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
		recovery: recovery.New(l, r),
	}
	y.MustRegister(result.sealed, result.reachable)
	result.ticker = ticker.New(
		constant.WatchInterval,
		result.Poll,
		result.recovery,
	)

	return result
}
