package service

func (s *Service) Place(name string) {
	s.record(s.store.Order(name).Name)
}
