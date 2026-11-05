package service

import "github.com/funtimecoding/soil/pkg/opnsense/request"

func (s *Service) AddHost(
	h *request.Host,
	apply bool,
) (string, error) {
	return guarded(
		s.guard,
		func() (string, error) {
			result, e := s.client.AddHost(h)

			if e != nil {
				return "", e
			}

			if f := s.reconfigureWhen(apply); f != nil {
				return "", f
			}

			return result, nil
		},
	)
}
