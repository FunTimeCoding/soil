package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.closeOnce.Do(func() { relational.CloseMapper(s.database) })
}
