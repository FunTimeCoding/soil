package peered

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.BandClient,
	g *peer.Guard,
) *Peered {
	return &Peered{client: c, guard: g}
}
