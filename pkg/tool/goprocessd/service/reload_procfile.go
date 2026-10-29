package service

func (s *Service) ReloadProcfile() error {
	return s.supervisor.ReloadProcfile()
}
