package reader

func (s *Reader) ClickRefresh() {
	if !s.Protocol.HasNodes(`button[aria-label="Refresh"]`) {
		return
	}

	s.Protocol.ClickQuery(`button[aria-label="Refresh"]`)
}
