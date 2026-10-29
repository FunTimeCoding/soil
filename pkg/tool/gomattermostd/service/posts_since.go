package service

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (s *Service) PostsSince(
	h *model.Channel,
	since time.Time,
) ([]*post.Post, error) {
	return guarded(
		s.guard,
		func() ([]*post.Post, error) {
			return s.mattermost.PostsSince(h, since)
		},
	)
}
