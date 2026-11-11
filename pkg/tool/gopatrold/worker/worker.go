package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/types/inspection"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

type Worker struct {
	round       func(context.Context) ([]*inspection.Inspection, []error)
	gauges      []*prometheus.GaugeVec
	reached     *prometheus.GaugeVec
	required    *prometheus.GaugeVec
	active      *prometheus.GaugeVec
	failed      *prometheus.GaugeVec
	rebuilding  *prometheus.GaugeVec
	passed      *prometheus.GaugeVec
	temperature *prometheus.GaugeVec
	warning     *prometheus.GaugeVec
	spare       *prometheus.GaugeVec
	used        *prometheus.GaugeVec
	media       *prometheus.GaugeVec
	clock       *worker.Worker
}
