package mock_service

func (s *Service) StartContainer(
	_ string,
	_ int,
	_ string,
) (string, error) {
	return "mock:ct-start", nil
}
