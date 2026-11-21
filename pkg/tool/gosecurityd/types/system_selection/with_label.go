package system_selection

func (s *Selection) WithLabel(
	key string,
	value string,
) *Selection {
	result := s.clone()

	if value == "" {
		delete(result.Labels, key)
	} else {
		result.Labels[key] = value
	}

	return result
}
