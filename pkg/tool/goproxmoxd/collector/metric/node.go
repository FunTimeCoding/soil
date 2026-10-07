package metric

import "github.com/prometheus/client_golang/prometheus"

type Node struct {
	Status         *prometheus.GaugeVec
	Processor      *prometheus.GaugeVec
	ProcessorCount *prometheus.GaugeVec
	MemoryUsed     *prometheus.GaugeVec
	MemoryTotal    *prometheus.GaugeVec
	DiskUsed       *prometheus.GaugeVec
	DiskTotal      *prometheus.GaugeVec
	Uptime         *prometheus.GaugeVec
	Version        *prometheus.GaugeVec
	UpdatePending  *prometheus.GaugeVec
}
