package service

func (s *Service) DeleteInternet(identifier int32) error {
	return s.attempt(
		func() error {
			return s.client.DeleteInternet(identifier)
		},
	)
}
