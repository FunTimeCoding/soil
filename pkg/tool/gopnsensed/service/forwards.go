package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/forward"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Forwards(phrase string) ([]*forward.Forward, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*forward.Forward, error) {
			return s.client.Forwards(phrase)
		},
	)
}
