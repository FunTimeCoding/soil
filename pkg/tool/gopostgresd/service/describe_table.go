package service

import (
	"context"
	"fmt"
)

func (s *Service) DescribeTable(
	x context.Context,
	instance string,
	schema string,
	table string,
) ([]map[string]any, error) {
	return s.Query(
		x,
		instance,
		fmt.Sprintf(describeTableSQL, schema, table),
	)
}
