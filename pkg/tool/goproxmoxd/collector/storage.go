package collector

import "github.com/prometheus/client_golang/prometheus"

type Storage struct {
	status *prometheus.GaugeVec
	used   *prometheus.GaugeVec
	total  *prometheus.GaugeVec
	shared *prometheus.GaugeVec
}
