package mock_service

func (s *Service) ShutdownMachine(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:shutdown", nil
}
