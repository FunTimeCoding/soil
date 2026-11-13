package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ThreadOf(identifier string) ([]*post.Post, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*post.Post, error) {
			p, e := s.mattermost.FindPost(identifier)

			if e != nil {
				return nil, e
			}

			return s.mattermost.Thread(p)
		},
	)
}
