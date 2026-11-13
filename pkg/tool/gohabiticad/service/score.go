package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/score"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Score(
	taskIdentifier string,
	direction string,
) (*score.Score, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*score.Score, error) {
			return s.client.Score(taskIdentifier, direction)
		},
	)
}
