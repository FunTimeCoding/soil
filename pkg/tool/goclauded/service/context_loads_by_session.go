package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"

func (s *Service) ContextLoadsBySession(
	sessionIdentifier string,
) ([]context_load.Load, error) {
	loads, e := s.store.ContextLoadsBySession(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	return s.redactLoads(loads), nil
}
