package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

func New(
	c face.PrometheusClient,
	g *peer.Guard,
) *Peered {
	return &Peered{client: c, guard: g}
}
