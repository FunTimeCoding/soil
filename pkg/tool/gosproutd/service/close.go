package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	relational.CloseMapper(s.database)
}
