package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/openbao/status"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Status() (*status.Status, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*status.Status, error) {
			return s.client.Health()
		},
	)
}
