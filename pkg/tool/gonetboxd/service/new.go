package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

func New(
	s *store.Store,
	c face.NetBoxClient,
	g *peer.Guard,
) *Service {
	return &Service{client: c, guard: g, store: s}
}
