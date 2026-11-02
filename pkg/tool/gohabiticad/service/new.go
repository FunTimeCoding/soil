package service

import (
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.HabiticaClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
