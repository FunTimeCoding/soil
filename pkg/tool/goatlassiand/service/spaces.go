package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Spaces() ([]*space.Space, error) {
	return peer.Guarded(
		context.Background(),
		s.confluenceGuard,
		func(context.Context) ([]*space.Space, error) {
			return s.confluence.Spaces()
		},
	)
}
