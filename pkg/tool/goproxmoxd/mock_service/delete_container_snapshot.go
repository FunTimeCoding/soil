package mock_service

func (s *Service) DeleteContainerSnapshot(
	_ string,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:ct-delete-snapshot", nil
}
