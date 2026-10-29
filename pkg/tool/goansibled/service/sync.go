package service

import "github.com/funtimecoding/soil/pkg/provision/types/update"

func (s *Service) Sync() (*update.Result, error) {
	return s.runner.Sync()
}
