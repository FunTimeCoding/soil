package service

func (s *Service) Prune() int64 {
	return s.store.Prune()
}
