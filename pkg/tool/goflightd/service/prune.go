package service

func (s *Service) Prune() []error {
	cutoff := s.clock().Add(-s.retention)
	events, e := s.store.PruneEventsBefore(cutoff)

	if e != nil {
		return []error{e}
	}

	snapshots, f := s.store.PruneSnapshotsBefore(cutoff)

	if f != nil {
		return []error{f}
	}

	if events > 0 || snapshots > 0 {
		s.logger.Structured("pruned", "events", events, "snapshots", snapshots)
	}

	return nil
}
