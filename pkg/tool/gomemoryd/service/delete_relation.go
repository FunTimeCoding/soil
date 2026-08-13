package service

func (s *Service) DeleteRelation(
	sourceIdentifier int64,
	targetIdentifier int64,
) (bool, error) {
	return s.store.DeleteRelation(sourceIdentifier, targetIdentifier)
}
