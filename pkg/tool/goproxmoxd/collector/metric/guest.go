package metric

import "github.com/prometheus/client_golang/prometheus"

type Guest struct {
	Status          *prometheus.GaugeVec
	Template        *prometheus.GaugeVec
	Tag             *prometheus.GaugeVec
	Processor       *prometheus.GaugeVec
	ProcessorCount  *prometheus.GaugeVec
	MemoryUsed      *prometheus.GaugeVec
	MemoryTotal     *prometheus.GaugeVec
	DiskUsed        *prometheus.GaugeVec
	DiskTotal       *prometheus.GaugeVec
	Uptime          *prometheus.GaugeVec
	NetworkReceive  *prometheus.GaugeVec
	NetworkTransmit *prometheus.GaugeVec
	DiskRead        *prometheus.GaugeVec
	DiskWritten     *prometheus.GaugeVec
}
