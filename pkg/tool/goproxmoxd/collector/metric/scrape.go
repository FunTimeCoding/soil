package metric

import "github.com/prometheus/client_golang/prometheus"

type Scrape struct {
	Success  *prometheus.GaugeVec
	Duration *prometheus.GaugeVec
}
