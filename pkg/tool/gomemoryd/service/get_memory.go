package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) GetMemory(identifier int64) (*record.Memory, error) {
	return s.store.GetMemory(identifier)
}
