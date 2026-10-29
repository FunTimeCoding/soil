package service

func (s *Service) ClearLog(name string) error {
	return s.supervisor.ClearLog(name)
}
