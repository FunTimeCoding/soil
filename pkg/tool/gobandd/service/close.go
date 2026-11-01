package service

func (s *Service) Close() {
	s.pool.Close()

	if s.netbox == nil {
		return
	}

	s.netboxGuard.Close()
	s.netbox.Close()
}
