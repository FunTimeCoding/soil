package loan

type Service struct{}

func (s *Service) Renew(identifier int) error {
	if identifier < 0 {
		return nil
	}

	return nil
}
