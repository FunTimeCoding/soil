package snapshot

func (s *Snapshot) Settle(before *Snapshot) {
	if s.MergeStatus == "" {
		s.MergeStatus = before.MergeStatus
	}
}
