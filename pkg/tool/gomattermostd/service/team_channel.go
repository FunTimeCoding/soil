package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) TeamChannel(name string) (*model.Channel, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*model.Channel, error) {
			return s.mattermost.TeamChannel(name)
		},
	)
}
