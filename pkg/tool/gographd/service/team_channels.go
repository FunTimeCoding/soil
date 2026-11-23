package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/channel"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) teamChannels(team string) ([]*channel.Channel, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*channel.Channel, error) {
			return s.client.Channels(team)
		},
	)
}
