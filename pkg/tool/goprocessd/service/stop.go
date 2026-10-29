package service

func (s *Service) Stop(name string) error {
	return s.supervisor.Stop(name)
}
