package collector

import "github.com/prometheus/client_golang/prometheus"

type Guest struct {
	status          *prometheus.GaugeVec
	template        *prometheus.GaugeVec
	tag             *prometheus.GaugeVec
	processor       *prometheus.GaugeVec
	processorCount  *prometheus.GaugeVec
	memoryUsed      *prometheus.GaugeVec
	memoryTotal     *prometheus.GaugeVec
	diskUsed        *prometheus.GaugeVec
	diskTotal       *prometheus.GaugeVec
	uptime          *prometheus.GaugeVec
	networkReceive  *prometheus.GaugeVec
	networkTransmit *prometheus.GaugeVec
	diskRead        *prometheus.GaugeVec
	diskWritten     *prometheus.GaugeVec
}
