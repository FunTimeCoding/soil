package service

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strings"
)

func (s *Service) SendDirectMessage(
	user string,
	text string,
	mention []string,
) (*message.Message, error) {
	if strings.TrimSpace(text) == "" {
		return nil, validation.New("chat message is empty")
	}

	h, e := s.CreateChat([]string{user}, "")

	if e != nil {
		return nil, e
	}

	return s.SendChatMessage(h.Identifier, text, mention)
}
