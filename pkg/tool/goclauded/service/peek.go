package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/peek"

func (s *Service) Peek(sessionIdentifier string) *peek.Peek {
	return s.client.Peek(sessionIdentifier)
}
