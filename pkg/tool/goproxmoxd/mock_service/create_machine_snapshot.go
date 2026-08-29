package mock_service

func (s *Service) CreateMachineSnapshot(
	_ string,
	_ int,
	_ string,
	_ string,
) (string, error) {
	return "mock:create-snapshot", nil
}
