package collector

import "github.com/prometheus/client_golang/prometheus"

func gauge(
	registry *prometheus.Registry,
	name string,
	help string,
	label []string,
) *prometheus.GaugeVec {
	result := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: name, Help: help},
		label,
	)
	registry.MustRegister(result)

	return result
}
