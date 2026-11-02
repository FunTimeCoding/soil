package closing_source

func (s *Source) Close() {
	s.closed++
}
