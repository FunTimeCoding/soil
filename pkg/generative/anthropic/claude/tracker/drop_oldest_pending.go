package tracker

func (s *State) dropOldestPending() {
	oldest := ""

	for identifier, c := range s.Pending {
		if oldest == "" || c.Timestamp < s.Pending[oldest].Timestamp {
			oldest = identifier
		}
	}

	delete(s.Pending, oldest)
}
