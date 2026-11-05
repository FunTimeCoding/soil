package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/sublime/face"
)

func New(
	c face.SublimeClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
