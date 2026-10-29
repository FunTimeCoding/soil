package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/habitica/gear"
)

func (s *Service) Equip(key string) (*gear.Gear, error) {
	if e := s.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return s.client.Equip(key)
		},
	); e != nil {
		return nil, e
	}

	return s.UserGear()
}
