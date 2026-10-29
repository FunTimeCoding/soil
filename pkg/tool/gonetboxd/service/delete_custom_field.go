package service

func (s *Service) DeleteCustomField(name string) error {
	return s.client.DeleteCustomField(name)
}
