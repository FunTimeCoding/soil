package service

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
)

func (s *Service) PostChannelMessage(
	team string,
	channel string,
	text string,
	mention []string,
) (*message.Message, error) {
	t, h, e := s.resolveChannel(team, channel)

	if e != nil {
		return nil, e
	}

	return s.post(destination.Channel(t, h), text, mention)
}
