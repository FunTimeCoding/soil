package service

func (s *Service) BeforeCommit(f func()) {
	s.beforeCommit = f
}
