package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service/check_result"

func (s *Service) Check(sessionIdentifier string) (*check_result.Result, error) {
	r, e := s.store.EnsureSession(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	result := check_result.New()
	result.Callsign = r.Callsign
	entries, e := s.store.DrainQueue(r.Callsign)

	if e != nil {
		return nil, e
	}

	result.Entries = entries
	result.Changed = len(entries) > 0
	s.notify()

	return result, nil
}
