package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) ChoiceSetByName(n string) (*choice_set.Set, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*choice_set.Set, error) {
			return s.client.ChoiceSet(n)
		},
	)
}
