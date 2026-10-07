package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	AlertsTotal  prometheus.Counter
	AlertsFiring prometheus.Gauge
	RecordsTotal prometheus.Gauge
	PollDuration prometheus.Histogram
	LastPollTime prometheus.Gauge
}
