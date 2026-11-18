package target

import (
	"github.com/funtimecoding/soil/pkg/nut/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	server string,
	c face.Client,
	g *peer.Guard,
) *Target {
	return &Target{Server: server, Client: c, Guard: g}
}
