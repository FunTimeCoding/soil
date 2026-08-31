package service

func (s *Service) recordFable(
	percent int,
	reset string,
) error {
	latest, e := s.store.LatestFableSnapshot()

	if e != nil {
		return e
	}

	if latest != nil && latest.Percent == percent && latest.Reset == reset {
		return nil
	}

	s.store.SaveFableSnapshot(percent, reset)
	s.store.TrimFableSnapshots()

	return nil
}
