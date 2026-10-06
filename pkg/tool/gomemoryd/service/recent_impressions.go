package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) RecentImpressions(since string) ([]record.Impression, error) {
	return s.store.RecentImpressions(since)
}
