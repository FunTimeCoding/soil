package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ChatHistory(
	chat string,
	limit int,
) ([]*message.Message, error) {
	if e := checkChat(chat); e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*message.Message, error) {
			return s.client.ChatMessages(chat, limit)
		},
	)
}
