package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) AddHost(
	h *request.Host,
	apply bool,
) (string, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (string, error) {
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
