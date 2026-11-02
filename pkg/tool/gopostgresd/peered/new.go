package peered

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/postgres/face"
)

func New(
	c face.PostgresClient,
	g *peer.Guard,
) *Peered {
	return &Peered{client: c, guard: g}
}
