package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/plate"

func (s *Service) CurrentPlate() *plate.Plate {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.current
}
