package service

import "github.com/funtimecoding/soil/pkg/opnsense/request"

func (s *Service) SetHost(
	identifier string,
	h *request.Host,
	apply bool,
) error {
	if e := s.source.SetHost(identifier, h); e != nil {
		return e
	}

	return s.reconfigureWhen(apply)
}
