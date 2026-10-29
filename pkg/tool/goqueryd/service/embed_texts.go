package service

import "context"

func (s *Service) embedTexts(
	x context.Context,
	texts []string,
) ([][]float32, error) {
	var result [][]float32
	e := s.guard.Attempt(
		x,
		func(y context.Context) error {
			var f error
			result, f = s.embedder.Embed(y, texts)

			return f
		},
	)

	return result, e
}
