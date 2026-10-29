package service

func (s *Service) Monitoring() bool {
	return s.monitor != nil
}
