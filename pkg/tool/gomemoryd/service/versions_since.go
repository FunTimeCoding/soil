package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) VersionsSince(
	since string,
	limit int,
	offset int,
) ([]store.Version, error) {
	return s.store.VersionsSince(since, limit, offset)
}
