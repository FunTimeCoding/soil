package service

func (s *Service) SetTabColor(
	sessionIdentifier string,
	red int,
	green int,
	blue int,
) error {
	return s.source.SetTabColor(sessionIdentifier, red, green, blue)
}
