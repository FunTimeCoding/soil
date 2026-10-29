package lease

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
)

func New(
	c face.Client,
	g *peer.Guard,
) *Collector {
	return &Collector{opnsense: c, guard: g}
}
