package startup

import "context"

func (s *Service) Rescan(x context.Context) []error {
	s.transcript.Sweep()
	s.transcript.PopulateCache()
	s.transcript.BackfillSessions()
	s.transcript.CheckConsistency()

	return s.search.CatchUpSearch(x)
}
