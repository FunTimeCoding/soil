package mock_service

func (s *Service) StartMachine(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:start", nil
}
