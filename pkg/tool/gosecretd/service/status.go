package service

import "github.com/funtimecoding/soil/pkg/openbao/status"

func (s *Service) Status() (*status.Status, error) {
	return s.source.Health()
}
