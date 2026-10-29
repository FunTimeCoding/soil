package kubernetes

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.Client,
	g *peer.Guard,
) *Collector {
	return &Collector{kubernetes: c, guard: g}
}
