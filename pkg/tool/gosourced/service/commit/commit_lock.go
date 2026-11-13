package commit

import "sync"

func (s *Service) CommitLock() sync.Locker {
	return &s.commits
}
