package service

func (s *Service) SearchProgress() (int64, int64) {
	return s.searchIndexed.Load(), s.searchTotal.Load()
}
