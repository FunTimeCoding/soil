package service

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Service) asQuery(text string) string {
	return join.Empty(s.model.QueryPrompt, text)
}
