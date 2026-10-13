package selection

func (s *Selection) Set(
	session string,
	instance string,
) {
	s.sessions.Store(session, instance)
}
