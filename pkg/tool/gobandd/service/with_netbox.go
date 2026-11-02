package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

func (s *Service) WithNetbox(
	n face.Client,
	g *peer.Guard,
) *Service {
	s.netbox = n
	s.netboxGuard = g

	return s
}
