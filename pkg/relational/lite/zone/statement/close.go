package statement

func (s *Statement) Close() error {
	return s.inner.Close()
}
