package backup

func (s *Service) Delete(
	instance string,
	name string,
) error {
	if _, e := s.find(instance); e != nil {
		return e
	}

	return s.store.Remove(instance, name)
}
