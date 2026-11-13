package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (s *Service) PostsSince(
	h *model.Channel,
	since time.Time,
) ([]*post.Post, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*post.Post, error) {
			return s.mattermost.Posts(h, since)
		},
	)
}
