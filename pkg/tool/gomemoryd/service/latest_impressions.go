package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) LatestImpressions(limit int) ([]record.Impression, error) {
	return s.store.LatestImpressions(limit)
}
