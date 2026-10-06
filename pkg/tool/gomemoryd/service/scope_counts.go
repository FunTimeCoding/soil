package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ScopeCounts() ([]record.ScopeCount, error) {
	return s.store.ScopeCounts()
}
