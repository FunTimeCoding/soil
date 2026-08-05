package service

import "time"

func (s *Service) sweepCallsignRelease() {
	cutoff := s.clock().Add(-72 * time.Hour)
	s.store.SweepCallsignRelease(cutoff)
}
