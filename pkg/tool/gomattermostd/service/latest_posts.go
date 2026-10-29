package service

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) LatestPosts(
	h *model.Channel,
	limit int,
) ([]*post.Post, error) {
	return guarded(
		s.guard,
		func() ([]*post.Post, error) {
			return s.mattermost.LatestPosts(h, limit)
		},
	)
}
