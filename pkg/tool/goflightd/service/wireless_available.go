package service

func (s *Service) WirelessAvailable() bool {
	_, e := s.wireless()

	return e == nil
}
