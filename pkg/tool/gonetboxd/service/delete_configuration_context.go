package service

func (s *Service) DeleteConfigurationContext(name string) error {
	return s.client.DeleteConfigurationContext(name)
}
