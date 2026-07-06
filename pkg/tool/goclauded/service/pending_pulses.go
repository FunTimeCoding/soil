package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Service) PendingPulses(sessionIdentifier string) ([]pulse.Pulse, error) {
	return s.store.PendingPulses(sessionIdentifier)
}
