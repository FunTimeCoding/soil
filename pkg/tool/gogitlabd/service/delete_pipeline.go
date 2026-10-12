package service

func (s *Service) DeletePipeline(
	project int64,
	identifier int64,
) error {
	return s.client.DeletePipeline(project, identifier)
}
