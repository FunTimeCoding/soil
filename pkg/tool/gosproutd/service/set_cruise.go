package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"time"
)

func (s *Service) SetCruise(
	session string,
	mode constant.Cruise,
	pace time.Duration,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.SetCruise(session, mode, pace)
	s.notifier.Notify()
}
