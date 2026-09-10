package service

func (s *Service) RunTimeoutSweep() {
	s.sweepInactivity()
	s.sweepCompleteTimeout()
	s.sweepCallsignRelease()
	s.cleanupQueue()
	s.sweepEmptySessions()
}
