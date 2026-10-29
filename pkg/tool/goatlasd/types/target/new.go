package target

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
)

func New(
	name string,
	c face.Client,
	g *peer.Guard,
) *Target {
	return &Target{Name: name, Client: c, Guard: g}
}
