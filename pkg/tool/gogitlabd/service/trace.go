package service

func (s *Service) Trace(
	instance string,
	project int64,
	jobIdentifier int64,
) (string, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return "", e
	}

	return c.Trace(project, jobIdentifier)
}
