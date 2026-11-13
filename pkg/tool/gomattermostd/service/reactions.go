package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) Reactions(p *model.Post) ([]*model.Reaction, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*model.Reaction, error) {
			return s.mattermost.Reactions(p)
		},
	)
}
