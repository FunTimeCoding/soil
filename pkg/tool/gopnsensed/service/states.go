package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/state"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) States(phrase string) ([]*state.State, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*state.State, error) {
			return s.client.States(phrase)
		},
	)
}
