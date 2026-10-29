package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) Users(
	page int,
	perPage int,
) ([]*model.User, error) {
	return guarded(
		s.guard,
		func() ([]*model.User, error) {
			return s.mattermost.Users(page, perPage)
		},
	)
}
