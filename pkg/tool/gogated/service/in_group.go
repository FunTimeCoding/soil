package service

func (s *Service) inGroup(account string) (bool, error) {
	return guarded(
		s.directoryGuard,
		func() (bool, error) {
			return s.directory.InGroup(account)
		},
	)
}
