package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) SearchUsers(
	query string,
	limit int,
) ([]*model.User, error) {
	return guarded(
		s.guard,
		func() ([]*model.User, error) {
			return s.mattermost.SearchUsers(
				s.mattermost.DefaultTeam(),
				query,
				limit,
			)
		},
	)
}
