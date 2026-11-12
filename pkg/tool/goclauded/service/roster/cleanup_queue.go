package roster

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) cleanupQueue() {
	s.store.CleanupQueue(s.clock().Add(-24 * time.Hour))
	s.store.DropAbandonedQueue(s.clock().Add(-constant.AbandonedQueueWindow))
}
