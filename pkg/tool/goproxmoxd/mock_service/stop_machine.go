package mock_service

func (s *Service) StopMachine(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:stop", nil
}
