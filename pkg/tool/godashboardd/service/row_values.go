package service

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout/row"

func (s *Service) rowValues(rows []*row.Row) ([]string, []error) {
	var result []string
	var failed []error

	for _, v := range rows {
		value, e := s.queryValue(v.Query)
		result = append(result, value)

		if e != nil {
			failed = append(failed, e)
		}
	}

	return result, failed
}
