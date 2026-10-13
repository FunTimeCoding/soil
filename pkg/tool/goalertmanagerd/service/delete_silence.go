package service

func (s *Service) DeleteSilence(
	instance string,
	identifier string,
) error {
	c, e := s.pool.Client(instance)

	if e != nil {
		return e
	}

	return c.DeleteSilence(identifier)
}
