package selection

func (s *Selection) Clear(session string) {
	s.sessions.Delete(session)
}
