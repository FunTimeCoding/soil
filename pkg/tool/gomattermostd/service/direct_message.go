package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) DirectMessage(
	u *model.User,
	text string,
) (*model.Post, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*model.Post, error) {
			return s.mattermost.DirectMessage(u, text)
		},
	)
}
