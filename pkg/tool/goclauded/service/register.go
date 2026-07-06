package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/ensure_result"
)

func (s *Service) Register(sessionIdentifier string) (*ensure_result.Result, error) {
	result, e := s.store.EnsureSession(sessionIdentifier)

	if e != nil {
		return nil, e
	}

	if result.Callsign == "" {
		return result, nil
	}

	if e := s.store.LogEvent(
		sessionIdentifier,
		constant.Register,
		result.Callsign,
		nil,
	); e != nil {
		return nil, e
	}

	s.notify()

	return result, nil
}
