package mock_service

func (s *Service) CreateContainerSnapshot(
	_ string,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:ct-create-snapshot", nil
}
