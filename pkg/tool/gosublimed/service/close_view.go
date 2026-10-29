package service

func (s *Service) CloseView(identifier int) error {
	return s.source.CloseView(identifier)
}
