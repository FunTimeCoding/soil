package service

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"

func (s *Service) Query(o *store.QueryOption) ([]store.UsageEvent, error) {
	return s.store.Recent(o)
}
