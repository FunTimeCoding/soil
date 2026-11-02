package peered

import "github.com/funtimecoding/soil/pkg/prometheus/face"

func (p *Peered) Prometheus() face.PrometheusClient {
	return p.prometheus
}
