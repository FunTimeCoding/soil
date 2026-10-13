package selection

func (s *Selection) Active(session string) (string, bool) {
	v, okay := s.sessions.Load(session)

	if !okay {
		return "", false
	}

	return v.(string), true
}
