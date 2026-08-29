package mock_service

func (s *Service) ShutdownContainer(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:ct-shutdown", nil
}
