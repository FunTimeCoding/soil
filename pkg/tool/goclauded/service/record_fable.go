package service

import "time"

func (s *Service) recordFable(
	percent int,
	reset string,
	resetAt *time.Time,
) error {
	latest, e := s.store.LatestFableSnapshot()

	if e != nil {
		return e
	}

	if latest != nil && latest.Percent == percent && latest.Reset == reset {
		return nil
	}

	s.store.SaveFableSnapshot(percent, reset, resetAt)
	s.store.TrimFableSnapshots()

	return nil
}
