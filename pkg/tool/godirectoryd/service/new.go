package service

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	d face.DirectoryClient,
	g *peer.Guard,
) *Service {
	return &Service{directory: d, guard: g}
}
