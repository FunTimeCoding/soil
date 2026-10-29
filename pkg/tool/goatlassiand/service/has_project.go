package service

func (s *Service) HasProject() bool {
	return len(s.project) > 0
}
