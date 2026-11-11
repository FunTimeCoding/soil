package worker

import "github.com/prometheus/client_golang/prometheus"

func gauge(
	y *prometheus.Registry,
	name string,
	help string,
	labels ...string,
) *prometheus.GaugeVec {
	result := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: name, Help: help},
		labels,
	)
	y.MustRegister(result)

	return result
}
