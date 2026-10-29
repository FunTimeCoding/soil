package service

func (s *Service) ReloadEnvironment() error {
	return s.supervisor.ReloadEnvironment()
}
