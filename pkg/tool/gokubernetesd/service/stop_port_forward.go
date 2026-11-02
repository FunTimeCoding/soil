package service

func (s *Service) StopPortForward(identifier string) error {
	return s.forwards.StopOne(identifier)
}
