package service

func (s *Service) Close() {
	s.client.Close()
}
