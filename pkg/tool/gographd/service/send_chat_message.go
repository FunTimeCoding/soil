package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"regexp"
	"strings"
)

func (s *Service) SendChatMessage(
	chat string,
	text string,
	mention []string,
) (*message.Message, error) {
	if e := checkChat(chat); e != nil {
		return nil, e
	}

	if strings.TrimSpace(text) == "" {
		return nil, validation.New("chat message is empty")
	}

	mentioned, e := s.resolveMembers(chat, mention)

	if e != nil {
		return nil, e
	}

	for i, m := range mentioned {
		text = regexp.MustCompile(
			join.Empty("(?i)@", regexp.QuoteMeta(mention[i])),
		).ReplaceAllLiteralString(text, join.Empty("@", m.Name))
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*message.Message, error) {
			return s.client.SendChatMessage(chat, text, mentioned)
		},
	)
}
