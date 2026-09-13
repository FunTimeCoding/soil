package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Service) LatestPulsesBySessions(
	sessionIdentifiers []string,
) (map[string]*pulse.Pulse, error) {
	return s.store.LatestPulsesBySessions(sessionIdentifiers)
}
