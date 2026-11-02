package service

func (s *Service) persist() error {
	e := s.save()

	if e != nil {
		s.discard()
	}

	return e
}
