package service

func (s *Service) Reveal(identifier string) (string, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	entry, _ := s.client.EntryByIdentifier(identifier)

	if entry == nil {
		return "", false
	}

	return entry.GetPassword(), true
}
