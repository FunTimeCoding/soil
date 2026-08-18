package service

func (s *Service) TagMembership() (map[string][]int64, error) {
	return s.store.TagMembership()
}
