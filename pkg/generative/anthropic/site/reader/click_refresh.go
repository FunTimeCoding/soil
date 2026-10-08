package reader

func (s *Reader) ClickRefresh() {
	if !s.Protocol.MustHasNodes(`button[aria-label="Refresh"]`) {
		return
	}

	s.Protocol.MustClickQuery(`button[aria-label="Refresh"]`)
}
