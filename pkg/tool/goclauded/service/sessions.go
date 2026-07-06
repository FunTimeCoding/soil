package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"

func (s *Service) Sessions() []*session.Session {
	return s.client.Sessions()
}
