package service

import "github.com/funtimecoding/soil/pkg/habitica/statistic"

func (s *Service) UserStats() (*statistic.Statistic, error) {
	return guarded(s.guard, s.client.UserStats)
}
