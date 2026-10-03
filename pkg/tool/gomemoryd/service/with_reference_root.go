package service

func (s *Service) WithReferenceRoot(root string) *Service {
	s.referenceRoot = root

	return s
}
