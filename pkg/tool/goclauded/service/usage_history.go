package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/usage_snapshot"
	"time"
)

func (s *Service) UsageHistory(since time.Duration) ([]usage_snapshot.Snapshot, error) {
	return s.store.UsageSnapshots(since)
}
