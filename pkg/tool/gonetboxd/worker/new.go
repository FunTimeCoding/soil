package worker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/intent"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	intents func() ([]*intent.Intent, error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	result := &Worker{intents: intents, metric: metric.New(y)}
	result.clock = worker.NewEvery(result.tick, l, r, interval)

	return result
}
