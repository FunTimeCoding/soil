package target

import (
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	host string,
	c face.PatrolClient,
	g *peer.Guard,
) *Target {
	return &Target{Host: host, Client: c, Guard: g}
}
