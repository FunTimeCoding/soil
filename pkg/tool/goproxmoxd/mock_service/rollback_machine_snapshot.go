package mock_service

func (s *Service) RollbackMachineSnapshot(
	_ string,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:rollback-snapshot", nil
}
