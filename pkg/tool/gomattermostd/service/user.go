package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) User(identifier string) (*model.User, error) {
	return guarded(
		s.guard,
		func() (*model.User, error) {
			return s.mattermost.User(identifier)
		},
	)
}
