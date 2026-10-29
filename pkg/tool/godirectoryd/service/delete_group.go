package service

func (s *Service) DeleteGroup(name string) error {
	return s.attempt(
		func() error { return s.directory.Delete(s.groupName(name)) },
	)
}
