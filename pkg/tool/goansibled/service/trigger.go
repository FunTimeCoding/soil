package service

import "github.com/funtimecoding/soil/pkg/provision/types/trigger"

func (s *Service) Trigger(request trigger.Request) error {
	return s.runner.Trigger(request)
}
