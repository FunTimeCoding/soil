package observation_tester

import "github.com/prometheus/client_golang/prometheus"

type Observation struct {
	Histogram *prometheus.HistogramVec
	registry  *prometheus.Registry
}
