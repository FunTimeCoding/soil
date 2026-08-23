package metric

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func (m *Metric) Exporter() http.Handler {
	return promhttp.HandlerFor(
		m.registry,
		promhttp.HandlerOpts{Registry: m.registry},
	)
}
