package service

func (s *Service) PollMonitor() []error {
	if e := s.RunMonitor(); e != nil {
		return []error{e}
	}

	return nil
}
