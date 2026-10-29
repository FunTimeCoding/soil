package process

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func New(
	c face.Client,
	g *peer.Guard,
	place string,
) *Collector {
	return &Collector{process: c, guard: g, place: place}
}
