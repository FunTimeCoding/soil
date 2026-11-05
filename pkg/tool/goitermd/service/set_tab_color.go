package service

func (s *Service) SetTabColor(
	sessionIdentifier string,
	red int,
	green int,
	blue int,
) error {
	return s.attempt(
		func() error {
			return s.client.SetTabColor(sessionIdentifier, red, green, blue)
		},
	)
}
