package service

func (s *Service) DeleteInternet(identifier int32) error {
	return s.client.DeleteInternet(identifier)
}
