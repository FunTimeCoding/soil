package label

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

func New(
	c face.Client,
	g *peer.Guard,
) *Collector {
	return &Collector{netbox: c, guard: g}
}
