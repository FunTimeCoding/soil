package service

func (s *Service) UserMail(identifier string) string {
	return s.store.UserMail(identifier)
}
