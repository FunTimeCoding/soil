package system

func (s *Service) Delete(name string) error {
	v, e := s.store.SystemByName(name)

	if e != nil {
		return e
	}

	if f := s.store.DeleteSystem(v); f != nil {
		return f
	}

	s.forget(name)

	return nil
}
