package service

func (s *Service) sessionName(identifier string) string {
	session, found, e := s.store.FindSession(identifier)

	if e != nil || !found {
		return ""
	}

	if session.Alias != nil && *session.Alias != "" {
		return *session.Alias
	}

	return session.Slug
}
