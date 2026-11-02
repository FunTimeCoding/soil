package transcript

func (s *Service) BackfillSessions() {
	for _, e := range s.store.UnenrichedSessions() {
		s.RefreshFromCache(e.Identifier)
	}
}
