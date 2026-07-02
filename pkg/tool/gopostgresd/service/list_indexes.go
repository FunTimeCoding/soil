package service

import (
	"context"
	"fmt"
)

func (s *Service) ListIndexes(
	x context.Context,
	instance string,
	schema string,
	table string,
) ([]map[string]any, error) {
	return s.Query(
		x,
		instance,
		fmt.Sprintf(listIndexesSQL, schema, table),
	)
}
