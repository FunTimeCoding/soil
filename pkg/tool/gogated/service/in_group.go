package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) inGroup(account string) (bool, error) {
	return peer.Guarded(
		context.Background(),
		s.directoryGuard,
		func(context.Context) (bool, error) {
			return s.directory.InGroup(account)
		},
	)
}
