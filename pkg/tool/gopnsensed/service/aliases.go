package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/alias"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Aliases(phrase string) ([]*alias.Alias, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*alias.Alias, error) {
			return s.client.Aliases(phrase)
		},
	)
}
