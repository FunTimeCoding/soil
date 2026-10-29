package service

func (s *Service) SaveView(
	identifier int,
	path string,
) error {
	return s.source.SaveView(identifier, path)
}
