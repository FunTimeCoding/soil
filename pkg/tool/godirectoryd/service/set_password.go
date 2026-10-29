package service

func (s *Service) SetPassword(
	account string,
	password string,
) error {
	return s.attempt(
		func() error {
			return s.directory.SetPassword(s.userName(account), password)
		},
	)
}
