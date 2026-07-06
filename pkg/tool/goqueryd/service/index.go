package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/index"

func (s *Service) Index(collection string) *index.Index {
	return s.store.Index(collection)
}
