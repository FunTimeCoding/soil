package alfa

func (s *Service) Hello() string {
	return s.Greet(s.Name)
}
