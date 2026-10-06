package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListRelations() ([]record.RelationOverview, error) {
	return s.store.ListRelations()
}
