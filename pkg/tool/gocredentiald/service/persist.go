package service

func (s *Service) persist() error {
	if !s.backedUp {
		target := s.client.Backup(s.clock())
		s.logger.Structured("backup written", "target", target)
		s.backedUp = true
	}

	return s.client.Save()
}
