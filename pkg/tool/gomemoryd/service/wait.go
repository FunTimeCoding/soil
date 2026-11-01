package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (s *Service) Wait(x context.Context) error {
	return peer.Await(
		x,
		s.logger,
		constant.Identity.Name(),
		s.query.Address(),
		s.query.Reachable,
	)
}
