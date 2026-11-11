package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/types/summary"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

type Worker struct {
	latest     func(context.Context) ([]*summary.Summary, error)
	status     *prometheus.GaugeVec
	completion *prometheus.GaugeVec
	size       *prometheus.GaugeVec
	repository *prometheus.GaugeVec
	damage     *prometheus.GaugeVec
	clock      *worker.Worker
}
