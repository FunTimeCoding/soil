package service

func (s *Service) UngroupTab(identifier int) error {
	return s.client.UngroupTab(identifier)
}
