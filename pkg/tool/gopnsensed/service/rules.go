package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/rule"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Rules(phrase string) ([]*rule.Rule, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*rule.Rule, error) {
			return s.client.Rules(phrase)
		},
	)
}
