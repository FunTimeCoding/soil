package service

func (s *Service) save() error {
	if !s.backedUp {
		target, e := s.client.Backup(s.clock())

		if e != nil {
			return e
		}

		s.logger.Structured("backup written", "target", target)
		s.backedUp = true
	}

	return s.client.Save()
}
