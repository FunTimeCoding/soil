package service

func (s *Service) DeviceTagNames(name string) ([]string, error) {
	return guarded(
		s.guard,
		func() ([]string, error) {
			return s.client.DeviceTagNames(name)
		},
	)
}
