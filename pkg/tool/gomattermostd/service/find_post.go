package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) FindPost(identifier string) (*model.Post, error) {
	return guarded(
		s.guard,
		func() (*model.Post, error) {
			return s.mattermost.FindPost(identifier)
		},
	)
}
