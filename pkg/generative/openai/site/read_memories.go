package site

func (s *Site) readMemories() string {
	if !s.session.HasNodes(`//table`) {
		return ""
	}

	s.session.WaitVisible(`//table//tbody/div[1]`)

	return s.session.Outer("table")
}
