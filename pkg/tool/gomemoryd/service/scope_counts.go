package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ScopeCounts() ([]store.ScopeCount, error) {
	return s.store.ScopeCounts()
}
