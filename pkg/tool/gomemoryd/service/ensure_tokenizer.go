package service

import "github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer"

func (s *Service) ensureTokenizer() error {
	if s.tokenizer != nil {
		return nil
	}

	e, f := tokenizer.New()

	if f != nil {
		return f
	}

	s.tokenizer = e

	return nil
}
