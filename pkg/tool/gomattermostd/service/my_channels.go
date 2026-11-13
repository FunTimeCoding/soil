package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) MyChannels() ([]*model.Channel, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*model.Channel, error) {
			me, e := s.mattermost.Me()

			if e != nil {
				return nil, e
			}

			return s.mattermost.Channels(s.mattermost.DefaultTeam(), me)
		},
	)
}
