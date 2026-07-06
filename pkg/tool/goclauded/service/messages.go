package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"

func (s *Service) Messages(sessionIdentifier string) []message.Message {
	return s.client.Messages(sessionIdentifier)
}
