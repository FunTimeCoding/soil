package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"

func (s *Service) ListSourceTypes() []record.SourceTypeTag {
	return s.store.ListSourceTypes()
}
