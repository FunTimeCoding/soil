package service

func (s *Service) DeleteList(identifier int) error {
	return s.attempt(func() error { return s.client.DeleteList(identifier) })
}
