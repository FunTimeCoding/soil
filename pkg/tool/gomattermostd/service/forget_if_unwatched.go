package service

func (s *Service) forgetIfUnwatched(root string) error {
	remaining, e := s.store.ByRoot(root)

	if e != nil {
		return e
	}

	if len(remaining) == 0 {
		s.watcher.Forget(root)
	}

	return nil
}
