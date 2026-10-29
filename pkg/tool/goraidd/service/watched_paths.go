package service

func (s *Service) WatchedPaths() []string {
	return []string{s.logCachePath, s.elitePath}
}
