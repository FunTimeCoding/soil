package service

func (s *Service) DeleteUser(account string) error {
	return s.attempt(
		func() error { return s.directory.Delete(s.userName(account)) },
	)
}
