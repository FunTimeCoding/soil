package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) SearchChannels(query string) ([]*model.Channel, error) {
	return guarded(
		s.guard,
		func() ([]*model.Channel, error) {
			return s.mattermost.SearchChannels(
				s.mattermost.DefaultTeam(),
				query,
			)
		},
	)
}
