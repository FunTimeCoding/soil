package service

func (s *Service) Run() error {
	return s.supervisor.Run()
}
