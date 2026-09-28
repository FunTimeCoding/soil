package service

func (s *Service) DirectoryConfigured() bool {
	return s.directory != nil
}
