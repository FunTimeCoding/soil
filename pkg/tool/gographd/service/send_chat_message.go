package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"strings"
)

func (s *Service) SendChatMessage(
	chat string,
	text string,
) (*message.Message, error) {
	if e := checkChat(chat); e != nil {
		return nil, e
	}

	if strings.TrimSpace(text) == "" {
		return nil, validation.New("chat message is empty")
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*message.Message, error) {
			return s.client.SendChatMessage(chat, text)
		},
	)
}
