package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Organizations() ([]response.Organization, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]response.Organization, error) {
			return s.client.Organizations()
		},
	)
}
