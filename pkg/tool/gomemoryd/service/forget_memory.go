package service

import "context"

func (s *Service) ForgetMemory(
	identifier int64,
	source string,
) error {
	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return e
	}

	if e = s.store.ForgetMemory(identifier, source); e != nil {
		return e
	}

	s.report(s.unindex(context.Background(), m.Scope, identifier))

	return nil
}
