package system

func (s *Service) Delete(name string) error {
	v, e := s.store.SystemByName(name)

	if e != nil {
		return e
	}

	return s.store.DeleteSystem(v)
}
