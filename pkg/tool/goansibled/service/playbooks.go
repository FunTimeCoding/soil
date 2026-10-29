package service

func (s *Service) Playbooks() []string {
	return s.runner.Playbooks()
}
