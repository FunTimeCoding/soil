package service

import "github.com/funtimecoding/soil/pkg/provision/model/run"

func (s *Service) Recent(limit int) ([]run.Run, error) {
	return s.store.Recent(limit)
}
