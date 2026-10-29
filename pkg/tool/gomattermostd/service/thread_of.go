package service

import "github.com/funtimecoding/soil/pkg/chat/mattermost/post"

func (s *Service) ThreadOf(identifier string) ([]*post.Post, error) {
	return guarded(
		s.guard,
		func() ([]*post.Post, error) {
			p, e := s.mattermost.FindPost(identifier)

			if e != nil {
				return nil, e
			}

			return s.mattermost.Thread(p)
		},
	)
}
