package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (s *Service) Keys() (*response.KeysReturn, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*response.KeysReturn, error) {
			return s.salt.Keys()
		},
	)
}
