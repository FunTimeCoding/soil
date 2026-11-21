package system_selection

func (s *Selection) WithSearch(search string) *Selection {
	result := s.clone()
	result.Search = search

	return result
}
