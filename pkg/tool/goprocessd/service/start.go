package service

func (s *Service) Start(name string) error {
	return s.supervisor.Start(name)
}
