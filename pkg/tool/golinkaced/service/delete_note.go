package service

func (s *Service) DeleteNote(identifier int) error {
	return s.attempt(func() error { return s.client.DeleteNote(identifier) })
}
