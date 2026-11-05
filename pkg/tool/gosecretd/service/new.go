package service

import (
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.OpenbaoClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
