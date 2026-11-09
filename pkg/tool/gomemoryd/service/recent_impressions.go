package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"time"
)

func (s *Service) RecentImpressions(span time.Duration) ([]record.Impression, error) {
	return s.store.RecentImpressions(windowStart(s.clock(), span))
}
