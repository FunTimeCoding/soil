package service

func (s *Service) Watching(root string) bool {
	return s.watcher != nil && s.watcher.WatchingRoot(root)
}
