package service

func (s *Service) Prepare() {
	s.Sweep()
	s.ClearBindings()
	s.PopulateCache()
	s.BackfillSessions()
	s.CheckConsistency()
}
