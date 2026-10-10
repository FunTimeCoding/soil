package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_row"
)

func (s *Service) Summary(o *summary_option.Option) ([]summary_row.Row, error) {
	return s.store.Summary(o)
}
