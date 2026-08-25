package collector

import "github.com/prometheus/client_golang/prometheus"

type Node struct {
	status         *prometheus.GaugeVec
	processor      *prometheus.GaugeVec
	processorCount *prometheus.GaugeVec
	memoryUsed     *prometheus.GaugeVec
	memoryTotal    *prometheus.GaugeVec
	diskUsed       *prometheus.GaugeVec
	diskTotal      *prometheus.GaugeVec
	uptime         *prometheus.GaugeVec
	version        *prometheus.GaugeVec
	updatePending  *prometheus.GaugeVec
}
