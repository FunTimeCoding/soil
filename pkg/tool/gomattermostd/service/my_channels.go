package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) MyChannels() ([]*model.Channel, error) {
	return guarded(
		s.guard,
		func() ([]*model.Channel, error) {
			me, e := s.mattermost.Me()

			if e != nil {
				return nil, e
			}

			return s.mattermost.Channels(s.mattermost.DefaultTeam(), me)
		},
	)
}
