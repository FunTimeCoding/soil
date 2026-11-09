package service

import (
	"github.com/funtimecoding/soil/pkg/time/window"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_row"
	"time"
)

func (s *Service) Summary(o *summary_option.Option) ([]summary_row.Row, error) {
	if o.Since == "" && o.Span > 0 {
		start, _ := window.Whole().Bound(s.clock(), o.Span).Start()
		o.Since = start.UTC().Format(time.RFC3339)
	}

	return s.store.Summary(o)
}
