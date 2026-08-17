package runner

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	runsTotal     *prometheus.CounterVec
	applyDuration prometheus.Histogram
	lastSuccess   prometheus.Gauge
}
