package service

func (s *Service) Trace(
	project int64,
	jobIdentifier int64,
) (string, error) {
	return s.client.Trace(project, jobIdentifier)
}
