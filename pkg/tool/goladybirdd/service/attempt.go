package service

import "context"

func (s *Service) attempt(call func() error) error {
	return s.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			return call()
		},
	)
}
