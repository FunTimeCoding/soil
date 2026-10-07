package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	RunsTotal     *prometheus.CounterVec
	ApplyDuration prometheus.Histogram
	LastSuccess   prometheus.Gauge
}
