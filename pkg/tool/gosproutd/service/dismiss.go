package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"

func (s *Service) Dismiss(identifier uint, state constant.State) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.Dismiss(identifier, state)
	s.notifier.Notify()
}
