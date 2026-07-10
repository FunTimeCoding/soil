package site

func (s *Site) ClickRefresh() {
	if !s.protocol.HasNodes(`button[aria-label="Refresh"]`) {
		return
	}

	s.protocol.ClickQuery(`button[aria-label="Refresh"]`)
}
