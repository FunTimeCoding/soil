package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store"

func (s *Service) ListSourceTypes() []store.SourceTypeTag {
	return s.store.ListSourceTypes()
}
