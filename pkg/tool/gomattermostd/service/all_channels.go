package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) AllChannels(
	limit int,
	offset int,
) ([]*model.ChannelWithTeamData, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*model.ChannelWithTeamData, error) {
			return s.mattermost.AllChannels(limit, offset)
		},
	)
}
