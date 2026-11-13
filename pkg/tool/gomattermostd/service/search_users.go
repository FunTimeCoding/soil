package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) SearchUsers(
	query string,
	limit int,
) ([]*model.User, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*model.User, error) {
			return s.mattermost.SearchUsers(
				s.mattermost.DefaultTeam(),
				query,
				limit,
			)
		},
	)
}
