package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) RecentVersions(
	since string,
	limit int,
) ([]record.Version, error) {
	return s.store.RecentVersions(since, limit, s.hiddenTag)
}
