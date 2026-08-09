package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ListRelations() ([]store.RelationOverview, error) {
	return s.store.ListRelations()
}
