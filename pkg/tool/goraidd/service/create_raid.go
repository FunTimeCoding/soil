package service

func (s *Service) CreateRaid(filenames []string) uint {
	return s.store.CreateRaid(filenames)
}
