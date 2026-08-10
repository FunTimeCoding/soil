package service

import "time"

func (s *Service) sweepCallsignRelease() {
	cutoff := s.clock().Add(-7 * 24 * time.Hour)
	s.store.SweepCallsignRelease(cutoff)
}
