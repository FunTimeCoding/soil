package target

import (
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	instance string,
	c face.PatrolClient,
	g *peer.Guard,
) *Target {
	return &Target{Instance: instance, Client: c, Guard: g}
}
