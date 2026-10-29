package service

func (s *Service) RestartAll() (int, error) {
	return s.supervisor.RestartWave()
}
