package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (s *Service) Minions() ([]response.Minion, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.Minion, error) {
			return s.salt.Minions()
		},
	)
}
