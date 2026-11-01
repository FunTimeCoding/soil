package reader

func (s *Reader) OpenWeek() {
	closed := `button[data-testid="pace-week-toggle"][aria-expanded="false"]`

	if !s.Protocol.MustHasNodes(closed) {
		return
	}

	s.Protocol.MustClickQuery(closed)
}
