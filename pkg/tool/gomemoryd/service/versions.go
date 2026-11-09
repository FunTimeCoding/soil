package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) Versions(
	since string,
	limit int,
	offset int,
) ([]record.Version, error) {
	return s.store.Versions(since, limit, offset)
}
