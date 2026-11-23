package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Chats(limit int) ([]*chat.Chat, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*chat.Chat, error) {
			return s.client.Chats(limit)
		},
	)
}
