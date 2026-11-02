package service

func (s *Service) reportReload(e error) {
	at := modified(s.path)

	if s.reported && at.Equal(s.reportedAt) {
		return
	}

	s.reported = true
	s.reportedAt = at
	s.logger.Structured("database reload failed", "error", e)
	s.reporter.CaptureException(e)
}
