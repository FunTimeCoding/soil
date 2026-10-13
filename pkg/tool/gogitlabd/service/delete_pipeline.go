package service

func (s *Service) DeletePipeline(
	instance string,
	project int64,
	identifier int64,
) error {
	c, e := s.pool.Client(instance)

	if e != nil {
		return e
	}

	return c.DeletePipeline(project, identifier)
}
