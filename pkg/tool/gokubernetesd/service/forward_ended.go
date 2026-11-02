package service

func (s *Service) forwardEnded(e error) {
	if e != nil {
		s.reporter.CaptureException(e)
	}
}
