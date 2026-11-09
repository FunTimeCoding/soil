package service

func (s *Service) WithPowerIntent() *Service {
	s.powerIntent = true

	return s
}
