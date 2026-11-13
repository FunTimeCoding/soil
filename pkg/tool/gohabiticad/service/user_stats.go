package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UserStats() (*statistic.Statistic, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*statistic.Statistic, error) {
			return s.client.UserStats()
		},
	)
}
