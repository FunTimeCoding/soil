package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SearchPosts(terms string) ([]*post.Post, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*post.Post, error) {
			list, e := s.mattermost.SearchPosts(
				s.mattermost.DefaultTeam(),
				terms,
			)

			if e != nil {
				return nil, e
			}

			result := post.NewSlice(post.FromList(list, true))

			if f := s.mattermost.Enrich(result); f != nil {
				return nil, f
			}

			return result, nil
		},
	)
}
