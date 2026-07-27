package service

func (s *Service) ForgetMemory(
	identifier int64,
	source string,
) error {
	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return e
	}

	if e = s.store.ForgetMemory(identifier, source); e != nil {
		return e
	}

	return s.indexer.Delete(
		ScopeCollection(m.Scope),
		memoryPath(identifier),
	)
}
