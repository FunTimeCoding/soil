package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.LinkAceClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
