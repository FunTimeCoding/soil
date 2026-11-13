package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/peer"
	"time"
)

func (s *Service) ChannelPosts(
	name string,
	since time.Time,
) ([]*post.Post, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*post.Post, error) {
			h, e := s.mattermost.TeamChannel(name)

			if e != nil {
				return nil, e
			}

			return s.mattermost.Posts(h, since)
		},
	)
}
