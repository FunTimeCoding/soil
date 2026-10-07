package service

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/summary_row"

func (s *Service) Summary(
	since string,
	until string,
	groupBy string,
) ([]summary_row.Row, error) {
	return s.store.Summary(since, until, groupBy)
}
