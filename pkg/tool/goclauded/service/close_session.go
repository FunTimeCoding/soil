package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) CloseSession(
	identifier string,
	reason string,
) error {
	r, found, e := s.store.FindSession(identifier)

	if e != nil {
		return e
	}

	if !found {
		return nil
	}

	if f := s.store.MarkClosed(identifier, reason); f != nil {
		return f
	}

	return s.store.LogEvent(
		identifier,
		constant.SessionEnd,
		r.CallsignValue(),
		map[string]string{"reason": reason},
	)
}
