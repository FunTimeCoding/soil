package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/intent"
	"github.com/funtimecoding/soil/pkg/worker"
)

type Worker struct {
	intents func() ([]*intent.Intent, error)
	metric  *metric.Metric
	clock   *worker.Worker
}
