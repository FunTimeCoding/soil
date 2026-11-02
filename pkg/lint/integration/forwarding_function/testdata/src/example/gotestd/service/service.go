package service

type Service struct{}

func (s *Service) Packages(limit int) []string {
	if limit < 0 {
		return nil
	}

	return nil
}
