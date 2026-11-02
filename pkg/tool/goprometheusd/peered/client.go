package peered

import "github.com/funtimecoding/soil/pkg/prometheus/face"

func (p *Peered) Client() face.PrometheusClient {
	return p.client
}
