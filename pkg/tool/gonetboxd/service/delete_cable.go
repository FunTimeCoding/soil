package service

func (s *Service) DeleteCable(identifier int32) error {
	return s.attempt(
		func() error {
			return s.client.DeleteCable(identifier)
		},
	)
}
