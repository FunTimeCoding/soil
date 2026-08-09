package service

func (s *Service) WithHiddenTag(tag string) *Service {
	s.hiddenTag = tag

	return s
}
