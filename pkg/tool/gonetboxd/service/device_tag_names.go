package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) DeviceTagNames(name string) ([]string, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]string, error) {
			return s.client.DeviceTagNames(name)
		},
	)
}
