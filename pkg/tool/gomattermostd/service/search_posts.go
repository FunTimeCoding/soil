package service

import "github.com/funtimecoding/soil/pkg/chat/mattermost/post"

func (s *Service) SearchPosts(terms string) ([]*post.Post, error) {
	return guarded(
		s.guard,
		func() ([]*post.Post, error) {
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
