package service

func (s *Service) DeleteUser(account string) error {
	return s.directory.Delete(s.userName(account))
}
