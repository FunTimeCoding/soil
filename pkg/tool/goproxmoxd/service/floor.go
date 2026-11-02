package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"

func (s *Service) Floor() *floor.Floor {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.floor == nil {
		return floor.New()
	}

	return s.floor
}
