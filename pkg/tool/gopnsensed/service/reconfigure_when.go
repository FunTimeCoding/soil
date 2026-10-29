package service

func (s *Service) reconfigureWhen(apply bool) error {
	if !apply {
		return nil
	}

	return s.source.ReconfigureDnsmasq()
}
