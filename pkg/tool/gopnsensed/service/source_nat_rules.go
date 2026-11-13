package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/source_nat"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SourceNatRules(phrase string) ([]*source_nat.Rule, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*source_nat.Rule, error) {
			return s.client.SourceNatRules(phrase)
		},
	)
}
