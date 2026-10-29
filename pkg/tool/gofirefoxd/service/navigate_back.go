package service

func (s *Service) NavigateBack(identifier int) error {
	return s.client.NavigateBack(identifier)
}
