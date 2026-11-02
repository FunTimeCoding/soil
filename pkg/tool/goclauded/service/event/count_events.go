package event

func (s *Service) CountEvents() (int64, error) {
	return s.store.CountEvents()
}
