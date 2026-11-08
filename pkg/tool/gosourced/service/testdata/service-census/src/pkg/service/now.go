package service

func (s *Service) Now() int {
	return stamp(s)
}
