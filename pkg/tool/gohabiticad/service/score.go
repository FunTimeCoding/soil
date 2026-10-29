package service

import "github.com/funtimecoding/soil/pkg/habitica/score"

func (s *Service) Score(
	taskIdentifier string,
	direction string,
) (*score.Score, error) {
	return guarded(
		s.guard,
		func() (*score.Score, error) {
			return s.client.Score(taskIdentifier, direction)
		},
	)
}
