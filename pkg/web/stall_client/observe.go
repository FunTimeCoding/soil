package stall_client

import "github.com/prometheus/client_golang/prometheus"

func (b *Builder) Observe(o prometheus.ObserverVec) *Builder {
	b.observer = o

	return b
}
