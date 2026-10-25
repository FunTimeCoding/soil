package statement

func (s *Statement) NumInput() int {
	return s.inner.NumInput()
}
