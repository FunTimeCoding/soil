package commit

func (s *Service) BeforeCommit(f func()) {
	s.beforeCommit = f
}
