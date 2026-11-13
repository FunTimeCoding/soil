package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Libraries() ([]*library.Library, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*library.Library, error) {
			return s.client.Libraries()
		},
	)
}
