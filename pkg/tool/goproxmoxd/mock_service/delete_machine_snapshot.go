package mock_service

func (s *Service) DeleteMachineSnapshot(
	_ string,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:delete-snapshot", nil
}
