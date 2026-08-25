package collector

import "github.com/prometheus/client_golang/prometheus"

type Scrape struct {
	success  *prometheus.GaugeVec
	duration *prometheus.GaugeVec
}
