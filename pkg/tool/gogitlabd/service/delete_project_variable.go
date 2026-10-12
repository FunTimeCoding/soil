package service

func (s *Service) DeleteProjectVariable(
	project int64,
	key string,
) error {
	return s.client.DeleteProjectVariable(project, key)
}
