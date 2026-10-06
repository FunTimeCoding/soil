package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListRelated(identifier int64) ([]record.Related, error) {
	return s.store.ListRelated(identifier)
}
