package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateChoiceSet(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*choice_set.Set, error) {
			return s.client.CreateChoiceSet(name, choices)
		},
	)
}
