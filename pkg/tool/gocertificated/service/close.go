package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/relational"
)

func (s *Service) Close() {
	s.guard.Close()

	if c, okay := s.forge.(face.Closer); okay {
		c.Close()
	}

	s.authorization.Close()
	relational.CloseMapper(s.database)
}
