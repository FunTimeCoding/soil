package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"strings"
)

func (s *Service) ChatHistory(
	chat string,
	limit int,
) ([]*message.Message, error) {
	if !constant.ChatPattern.MatchString(chat) ||
		strings.Contains(chat, constant.ParentStep) {
		return nil, validation.New(
			"chat identifier has an unexpected shape: %s",
			chat,
		)
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*message.Message, error) {
			return s.client.ChatMessages(chat, limit)
		},
	)
}
