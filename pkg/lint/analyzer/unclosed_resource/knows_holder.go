package unclosed_resource

func (s *Summaries) knowsHolder(holder string) bool {
	if s == nil {
		return false
	}

	_, result := s.releases[holder]

	return result
}
