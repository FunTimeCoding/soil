package service

func (s *Service) Delete(identifier uint) error {
	if e := s.store.Delete(identifier); e != nil {
		return e
	}

	s.notifier.Notify()

	return nil
}
