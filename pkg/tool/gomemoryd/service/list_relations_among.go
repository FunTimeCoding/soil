package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListRelationsAmong(
	identifiers []int64,
) ([]record.RelationOverview, error) {
	return s.store.ListRelationsAmong(identifiers)
}
