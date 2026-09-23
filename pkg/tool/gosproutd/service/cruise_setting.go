package service

import "github.com/funtimecoding/soil/pkg/tool/gosproutd/store/cruise"

func (s *Service) CruiseSetting(session string) *cruise.Cruise {
	return s.store.CruiseSetting(session)
}
