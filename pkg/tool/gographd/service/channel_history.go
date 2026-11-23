package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ChannelHistory(
	team string,
	channel string,
	limit int,
) ([]*message.Message, error) {
	t, h, e := s.resolveChannel(team, channel)

	if e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*message.Message, error) {
			return s.client.ChannelMessages(t, h, limit)
		},
	)
}
