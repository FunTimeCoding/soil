package cruise

import "github.com/funtimecoding/soil/pkg/tool/goclauded/cruise_reading"

func (s *Service) wakes(sessionIdentifier string) (bool, error) {
	v, found, e := s.store.FindSession(sessionIdentifier)

	if e != nil || !found {
		return false, e
	}

	setting, f := s.store.CruiseSetting(sessionIdentifier)

	if f != nil {
		return false, f
	}

	return cruise_reading.Read(v, s.Grace(setting)).Wakes(s.clock()), nil
}
