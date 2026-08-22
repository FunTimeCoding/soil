package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Service) FindLatestPulse(
	sessionIdentifier string,
) (*pulse.Pulse, bool, error) {
	return s.store.FindLatestPulse(sessionIdentifier)
}
