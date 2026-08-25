package service

func (s *Service) BindModelContextSession(
	name string,
	modelContextSessionIdentifier string,
) error {
	if e := s.store.BindModelContextSession(
		name,
		modelContextSessionIdentifier,
	); e != nil {
		return e
	}

	s.notify()

	return nil
}
