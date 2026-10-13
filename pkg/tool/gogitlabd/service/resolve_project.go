package service

func (s *Service) ResolveProject(
	instance string,
	identifier string,
) (int64, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return 0, e
	}

	return c.ResolveProject(identifier)
}
