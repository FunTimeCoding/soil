package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) GetMemoryHistory(identifier int64) ([]record.Version, error) {
	return s.store.GetMemoryHistory(identifier)
}
