package system_selection

func (s *Selection) WithMinimum(m float64) *Selection {
	result := s.clone()
	result.Minimum = m

	return result
}
