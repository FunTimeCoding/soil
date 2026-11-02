package roster

import "time"

func (s *Service) cleanupQueue() {
	s.store.CleanupQueue(s.clock().Add(-24 * time.Hour))
}
