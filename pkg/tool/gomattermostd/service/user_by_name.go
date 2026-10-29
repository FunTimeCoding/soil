package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) UserByName(name string) (*model.User, error) {
	return guarded(
		s.guard,
		func() (*model.User, error) {
			return s.mattermost.UserByName(name)
		},
	)
}
