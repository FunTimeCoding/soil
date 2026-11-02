package service

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"

func (s *Service) storeFloor(f *floor.Floor) {
	s.mutex.Lock()
	previous := s.floor
	s.floor = f
	s.mutex.Unlock()

	if previous == nil || !previous.Equal(*f) {
		s.notifier.Notify()
	}
}
