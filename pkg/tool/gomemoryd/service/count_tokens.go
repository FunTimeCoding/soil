package service

import "github.com/funtimecoding/soil/pkg/notation"

func (s *Service) countTokens(v any) int {
	return s.tokenizer.Count(notation.MarshalIndent(v))
}
