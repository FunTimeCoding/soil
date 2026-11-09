package worker

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/metric"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/outcome"
	"github.com/funtimecoding/soil/pkg/worker"
)

type Worker struct {
	record  func() (*outcome.Outcome, error)
	metrics *metric.Metric
	logger  *logger.Logger
	clock   *worker.Worker
}
