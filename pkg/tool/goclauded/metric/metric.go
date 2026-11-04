package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	reconciled *prometheus.CounterVec
	indexed    *prometheus.GaugeVec
	delivered  prometheus.Histogram
	overBudget prometheus.Counter
}
