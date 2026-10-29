package service

import "github.com/funtimecoding/soil/pkg/habitica/statistic"

func (s *Service) Allocate(stat string) (*statistic.Statistic, error) {
	return guarded(
		s.guard,
		func() (*statistic.Statistic, error) {
			return s.client.Allocate(stat)
		},
	)
}
