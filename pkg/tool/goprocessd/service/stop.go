package service

func (s *Service) Stop(name string) error {
	return s.supervisor.StopProcess(name)
}
