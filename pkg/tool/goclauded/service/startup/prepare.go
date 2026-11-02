package startup

func (s *Service) Prepare() {
	s.transcript.Sweep()
	s.roster.ClearBindings()
	s.transcript.PopulateCache()
	s.transcript.BackfillSessions()
	s.transcript.CheckConsistency()
}
