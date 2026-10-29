package service

import "github.com/funtimecoding/soil/pkg/provision/model/run"

func (s *Service) Find(identifier uint) (*run.Run, error) {
	return s.store.Find(identifier)
}
