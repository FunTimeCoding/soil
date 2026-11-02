package roster

func (s *Service) RunTimeoutSweep() []error {
	result := s.sweepInactivity()
	result = append(result, s.sweepCompleteTimeout()...)
	s.sweepCallsignRelease()
	s.cleanupQueue()

	return append(result, s.sweepEmptySessions()...)
}
