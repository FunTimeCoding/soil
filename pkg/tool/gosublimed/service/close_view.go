package service

func (s *Service) CloseView(identifier int) error {
	return s.attempt(
		func() error {
			return s.client.CloseView(identifier)
		},
	)
}
