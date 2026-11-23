package service

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.TeamsClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
