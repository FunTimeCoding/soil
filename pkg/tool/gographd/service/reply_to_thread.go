package service

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
)

func (s *Service) ReplyToThread(
	team string,
	channel string,
	identifier string,
	text string,
	mention []string,
) (*message.Message, error) {
	if !constant.MessagePattern.MatchString(identifier) {
		return nil, validation.New(
			"message identifier is not a number: %s",
			identifier,
		)
	}

	t, h, e := s.resolveChannel(team, channel)

	if e != nil {
		return nil, e
	}

	return s.post(destination.Thread(t, h, identifier), text, mention)
}
