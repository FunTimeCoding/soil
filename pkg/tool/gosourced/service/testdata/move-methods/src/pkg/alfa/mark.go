package alfa

func (s *Service) Mark() string {
	return s.Tag(s.Name)
}
