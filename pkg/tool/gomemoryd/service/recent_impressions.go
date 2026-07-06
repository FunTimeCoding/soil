package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) RecentImpressions(since string) ([]store.Impression, error) {
	return s.store.RecentImpressions(since)
}
