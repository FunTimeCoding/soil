package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CustomFields() ([]*custom_field.Field, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*custom_field.Field, error) {
			return s.client.CustomFields()
		},
	)
}
