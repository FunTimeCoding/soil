package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/types/poll"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

type Worker struct {
	read        func() *poll.Poll
	measures    map[string]*prometheus.GaugeVec
	status      *prometheus.GaugeVec
	information *prometheus.GaugeVec
	attached    *prometheus.GaugeVec
	reachable   *prometheus.GaugeVec
	clock       *worker.Worker
}
