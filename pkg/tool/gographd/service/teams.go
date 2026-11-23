package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Teams() ([]*team.Team, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*team.Team, error) {
			return s.client.Teams()
		},
	)
}
