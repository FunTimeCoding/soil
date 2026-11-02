package peered

import "github.com/funtimecoding/soil/pkg/peer"

func (p *Peered) PrometheusGuard() *peer.Guard {
	return p.prometheusGuard
}
