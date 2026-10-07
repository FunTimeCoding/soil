package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
)

func (s *Service) Query(o *query_option.Option) ([]usage_event.Event, error) {
	return s.store.Recent(o)
}
