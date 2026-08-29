package mock_service

func (s *Service) ResetMachine(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:reset", nil
}
