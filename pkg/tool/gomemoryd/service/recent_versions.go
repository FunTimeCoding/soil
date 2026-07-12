package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) RecentVersions(
	since string,
	limit int,
) ([]store.Version, error) {
	return s.store.RecentVersions(since, limit)
}
