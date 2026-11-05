package service

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.JellyfinClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g}
}
