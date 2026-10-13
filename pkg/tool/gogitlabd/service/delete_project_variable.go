package service

func (s *Service) DeleteProjectVariable(
	instance string,
	project int64,
	key string,
) error {
	c, e := s.pool.Client(instance)

	if e != nil {
		return e
	}

	return c.DeleteProjectVariable(project, key)
}
