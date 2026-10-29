package service

func (s *Service) File(identifier string) ([]byte, error) {
	return guarded(
		s.guard,
		func() ([]byte, error) {
			return s.mattermost.File(identifier)
		},
	)
}
