package service

func (s *Service) DeviceTagNames(name string) ([]string, error) {
	return s.client.DeviceTagNames(name)
}
