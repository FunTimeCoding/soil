package alfa

func (s *Service) Shout() string {
	return s.Loud(s.Name)
}
