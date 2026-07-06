package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Service) LatestPulse(sessionIdentifier string) (*pulse.Pulse, error) {
	return s.store.LatestPulse(sessionIdentifier)
}
