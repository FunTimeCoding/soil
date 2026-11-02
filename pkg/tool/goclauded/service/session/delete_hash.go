package session

func (s *Service) DeleteHash(identifier string) string {
	r, found, e := s.store.FindSession(identifier)

	if e != nil || !found {
		return ""
	}

	return r.DeleteHash()
}
