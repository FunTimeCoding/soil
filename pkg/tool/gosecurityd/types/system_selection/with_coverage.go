package system_selection

func (s *Selection) WithCoverage(show string) *Selection {
	result := s.clone()
	result.Show = show

	return result
}
