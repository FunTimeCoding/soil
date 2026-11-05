package service

func (s *Service) DeleteChoiceSet(name string) error {
	return s.attempt(
		func() error {
			return s.client.DeleteChoiceSet(name)
		},
	)
}
