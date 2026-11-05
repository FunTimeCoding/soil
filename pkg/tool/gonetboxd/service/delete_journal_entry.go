package service

func (s *Service) DeleteJournalEntry(identifier int32) error {
	return s.attempt(
		func() error {
			return s.client.DeleteJournalEntry(identifier)
		},
	)
}
