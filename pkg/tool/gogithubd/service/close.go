package service

import "github.com/funtimecoding/soil/pkg/face"

func (s *Service) Close() {
	if c, okay := s.source.(face.Closer); okay {
		c.Close()
	}
}
