package service

func (s *Service) Count() (int, error) {
	return s.store.Count()
}
