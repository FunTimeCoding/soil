package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/worker"
)

type Worker struct {
	service *service.Service
	metric  *metric.Metric
	clock   *worker.Worker
}
