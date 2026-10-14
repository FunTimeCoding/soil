package unclosed_resource

func (s *Summaries) releasesField(
	holder string,
	field string,
) bool {
	if s == nil {
		return false
	}

	return s.releases[holder][field]
}
