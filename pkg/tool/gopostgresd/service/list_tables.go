package service

import (
	"context"
	"fmt"
)

func (s *Service) ListTables(
	x context.Context,
	instance string,
	schema string,
) ([]map[string]any, error) {
	return s.Query(
		x,
		instance,
		fmt.Sprintf(listTablesSQL, schema),
	)
}
