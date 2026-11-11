package service

import (
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.BackrestClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
