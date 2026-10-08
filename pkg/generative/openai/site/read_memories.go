package site

func (s *Site) readMemories() string {
	if !s.session.MustHasNodes(`//table`) {
		return ""
	}

	s.session.MustWaitVisible(`//table//tbody/div[1]`)

	return s.session.MustOuter("table")
}
