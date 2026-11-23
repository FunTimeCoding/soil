package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
)

func (s *Service) ThreadReplies(
	team string,
	channel string,
	identifier string,
	limit int,
) ([]*message.Message, error) {
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

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*message.Message, error) {
			return s.client.Replies(t, h, identifier, limit)
		},
	)
}
