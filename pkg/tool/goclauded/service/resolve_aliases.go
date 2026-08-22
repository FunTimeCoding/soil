package service

func (s *Service) resolveAliases(sessionIDs map[string]bool) map[string]string {
	result := map[string]string{}

	for identifier := range sessionIDs {
		session, found, e := s.store.FindSession(identifier)

		if e != nil || !found {
			continue
		}

		if session.Alias != nil && *session.Alias != "" {
			result[identifier] = *session.Alias
		}
	}

	return result
}
