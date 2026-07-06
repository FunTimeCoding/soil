package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Service) PulsesBySession(sessionIdentifier string) ([]pulse.Pulse, error) {
	return s.store.PulsesBySession(sessionIdentifier)
}
