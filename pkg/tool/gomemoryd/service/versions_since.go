package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) VersionsSince(
	since string,
	limit int,
	offset int,
) ([]record.Version, error) {
	return s.store.VersionsSince(since, limit, offset)
}
