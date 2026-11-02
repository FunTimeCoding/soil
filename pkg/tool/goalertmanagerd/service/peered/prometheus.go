package peered

import "github.com/funtimecoding/soil/pkg/prometheus/face"

func (p *Peered) Prometheus() face.Client {
	return p.prometheus
}
