package service

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

func (s *Service) rowValues(rows []*board.Row) []string {
	var result []string

	for _, v := range rows {
		result = append(result, s.queryValue(v.Query))
	}

	return result
}
