package mock_service

func (s *Service) StopContainer(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:ct-stop", nil
}
