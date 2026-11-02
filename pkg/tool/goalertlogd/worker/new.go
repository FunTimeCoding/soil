package worker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/metric"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/outcome"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	record func(time.Time) (*outcome.Outcome, error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	result := &Worker{record: record, metrics: metric.New(y), logger: l}
	result.clock = worker.NewEvery(result.tick, l, r, interval)

	return result
}
