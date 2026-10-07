package metric

import "github.com/prometheus/client_golang/prometheus"

type Storage struct {
	Status *prometheus.GaugeVec
	Used   *prometheus.GaugeVec
	Total  *prometheus.GaugeVec
	Shared *prometheus.GaugeVec
}
