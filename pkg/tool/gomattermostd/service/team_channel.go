package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) TeamChannel(name string) (*model.Channel, error) {
	return guarded(
		s.guard,
		func() (*model.Channel, error) {
			return s.mattermost.TeamChannel(name)
		},
	)
}
