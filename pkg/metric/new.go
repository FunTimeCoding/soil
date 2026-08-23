package metric

import "github.com/prometheus/client_golang/prometheus"

func New() *Metric {
	return &Metric{registry: prometheus.NewRegistry()}
}
