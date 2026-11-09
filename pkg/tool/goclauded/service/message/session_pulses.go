package message

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"

func (s *Service) SessionPulses(sessionIdentifier string) ([]pulse.Pulse, error) {
	return s.store.SessionPulses(sessionIdentifier)
}
