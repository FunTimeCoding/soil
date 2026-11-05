package service

import "github.com/funtimecoding/soil/pkg/openbao/status"

func (s *Service) Status() (*status.Status, error) {
	return guarded(s.guard, s.client.Health)
}
