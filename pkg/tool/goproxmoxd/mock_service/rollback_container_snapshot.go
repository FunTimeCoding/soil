package mock_service

func (s *Service) RollbackContainerSnapshot(
	_ string,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:ct-rollback-snapshot", nil
}
