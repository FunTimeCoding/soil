package service

import "github.com/mattermost/mattermost/server/public/model"

func (s *Service) AllChannels(
	limit int,
	offset int,
) ([]*model.ChannelWithTeamData, error) {
	return guarded(
		s.guard,
		func() ([]*model.ChannelWithTeamData, error) {
			return s.mattermost.AllChannels(limit, offset)
		},
	)
}
