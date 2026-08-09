package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ListRelationsAmong(
	identifiers []int64,
) ([]store.RelationOverview, error) {
	return s.store.ListRelationsAmong(identifiers)
}
