package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) LatestPosts(
	h *model.Channel,
	limit int,
) ([]*post.Post, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*post.Post, error) {
			return s.mattermost.RecentPosts(h, limit)
		},
	)
}
