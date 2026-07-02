package service

import (
	"context"
	"fmt"
)

func (s *Service) Query(
	x context.Context,
	instance string,
	sql string,
) ([]map[string]any, error) {
	p, e := s.Pool(instance)

	if e != nil {
		return nil, e
	}

	rows, e := p.Query(x, sql)

	if e != nil {
		return nil, fmt.Errorf("query: %w", e)
	}

	defer rows.Close()

	return scanRows(rows)
}
