package service

func (s *Service) IsAdministrator(subject string) (bool, error) {
	if subject == "" {
		return false, nil
	}

	mail, e := s.store.UserMail(subject)

	if e != nil {
		return false, e
	}

	return mail == s.superUserMail, nil
}
