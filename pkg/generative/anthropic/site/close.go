package site

func (s *Site) Close() {
	s.session.Close()
}
