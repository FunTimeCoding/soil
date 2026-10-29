package service

func (s *Service) DeleteLink(identifier int) error {
	return s.attempt(func() error { return s.client.DeleteLink(identifier) })
}
