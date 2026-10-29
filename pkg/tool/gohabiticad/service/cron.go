package service

import "github.com/funtimecoding/soil/pkg/habitica/cron"

func (s *Service) Cron() (*cron.Cron, error) {
	return guarded(s.guard, s.client.Cron)
}
