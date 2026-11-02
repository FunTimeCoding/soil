package document

func (s *Service) backfillCompletionSequences() {
	completions := s.store.ListCompletions()
	sessions := map[string]int{}

	for _, c := range completions {
		if c.Sequence > 0 {
			if c.Sequence > sessions[c.SessionIdentifier] {
				sessions[c.SessionIdentifier] = c.Sequence
			}

			continue
		}

		sessions[c.SessionIdentifier]++
		s.store.UpdateCompletionSequence(
			c.Identifier,
			sessions[c.SessionIdentifier],
		)
	}
}
