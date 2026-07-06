package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func (s *Service) ListContexts() []result.ContextEntry {
	return s.store.ListContexts()
}
