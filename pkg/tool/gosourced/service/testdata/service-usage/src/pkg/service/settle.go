package service

import "example/pkg/store"

func (s *Service) Settle() error {
	return s.store.Atomically(
		func(t *store.Store) error {
			return s.step(t)
		},
	)
}
