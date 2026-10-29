package service

func (s *Service) Restart(name string) error {
	return s.supervisor.Restart(name)
}
