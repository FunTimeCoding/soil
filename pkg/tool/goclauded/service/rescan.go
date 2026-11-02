package service

import "context"

func (s *Service) Rescan(x context.Context) []error {
	s.Sweep()
	s.PopulateCache()
	s.BackfillSessions()
	s.CheckConsistency()

	return s.CatchUpSearch(x)
}
