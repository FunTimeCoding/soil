package service

func (s *Service) ReportPath(name string) (string, error) {
	return s.reportFile(name)
}
