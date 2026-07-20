package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
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
		fmt.Sprintf(constant.ListIndexesQuery, schema, table),
	)
}
