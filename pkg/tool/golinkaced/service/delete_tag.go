package service

func (s *Service) DeleteTag(identifier int) error {
	return s.attempt(func() error { return s.client.DeleteTag(identifier) })
}
