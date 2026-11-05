package service

func (s *Service) DeleteConfigurationContext(name string) error {
	return s.attempt(
		func() error {
			return s.client.DeleteConfigurationContext(name)
		},
	)
}
