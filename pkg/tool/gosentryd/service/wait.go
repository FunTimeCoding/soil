package service

import "context"

func (s *Service) Wait(x context.Context) error {
	return s.guard.Wait(x)
}
