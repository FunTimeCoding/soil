package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) LatestImpressions(limit int) ([]store.Impression, error) {
	return s.store.LatestImpressions(limit)
}
