package worker

import "github.com/prometheus/client_golang/prometheus"

func (w *Worker) gauge(
	y *prometheus.Registry,
	name string,
	help string,
	labels []string,
) *prometheus.GaugeVec {
	result := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: name, Help: help},
		labels,
	)
	y.MustRegister(result)
	w.gauges = append(w.gauges, result)

	return result
}
