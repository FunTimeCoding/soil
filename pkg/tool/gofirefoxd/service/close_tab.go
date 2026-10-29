package service

func (s *Service) CloseTab(identifier int) error {
	return s.client.CloseTab(identifier)
}
