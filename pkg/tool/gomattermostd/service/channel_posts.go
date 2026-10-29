package service

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"time"
)

func (s *Service) ChannelPosts(
	name string,
	since time.Time,
) ([]*post.Post, error) {
	return guarded(
		s.guard,
		func() ([]*post.Post, error) {
			h, e := s.mattermost.TeamChannel(name)

			if e != nil {
				return nil, e
			}

			return s.mattermost.PostsSince(h, since)
		},
	)
}
