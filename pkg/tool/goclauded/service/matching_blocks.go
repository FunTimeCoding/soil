package service

func (s *Service) MatchingBlocks(
	session string,
	query string,
	kinds []string,
) (map[string]bool, error) {
	return s.search.MatchingBlocks(session, query, kinds)
}
