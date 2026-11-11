package alfa

func (s *Service) Repeat() string {
	return s.Echo(s.Name)
}
