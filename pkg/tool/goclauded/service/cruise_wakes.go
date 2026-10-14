package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/cruise"

func (s *Service) cruiseWakes(sessionIdentifier string) (bool, error) {
	v, found, e := s.store.FindSession(sessionIdentifier)

	if e != nil || !found {
		return false, e
	}

	setting, f := s.store.CruiseSetting(sessionIdentifier)

	if f != nil {
		return false, f
	}

	return cruise.Wakes(cruise.Read(v, setting), s.clock()), nil
}
