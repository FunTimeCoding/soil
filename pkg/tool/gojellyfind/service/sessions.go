package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Sessions() ([]*session.Session, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*session.Session, error) {
			return s.client.Sessions()
		},
	)
}
