package service

func (s *Service) Tally() int {
	return s.store.Tally().Count
}
