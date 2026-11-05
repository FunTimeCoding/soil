package service

func (s *Service) SaveView(
	identifier int,
	path string,
) error {
	return s.attempt(
		func() error {
			return s.client.SaveView(identifier, path)
		},
	)
}
