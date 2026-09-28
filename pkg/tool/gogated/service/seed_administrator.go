package service

func (s *Service) SeedAdministrator(
	mail string,
	password string,
) error {
	return s.store.SeedAdministrator(mail, password)
}
