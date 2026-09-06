package scan

func (s *Service) guarded() bool {
	return s.ModelContext || s.Generated || s.Web
}
