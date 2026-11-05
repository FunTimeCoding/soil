package service

import (
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.OpnsenseClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
