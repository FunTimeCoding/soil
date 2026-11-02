package peered

import "github.com/funtimecoding/soil/pkg/prometheus/face"

func (p *Peered) Client() face.Client {
	return p.client
}
