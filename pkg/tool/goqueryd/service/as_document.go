package service

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Service) asDocument(text string) string {
	return join.Empty(s.model.DocumentPrompt, text)
}
