package service

import (
	"context"
	"fmt"
)

func (s *Service) TableSizes(
	x context.Context,
	instance string,
	schema string,
) ([]map[string]any, error) {
	return s.Query(
		x,
		instance,
		fmt.Sprintf(tableSizesSQL, schema),
	)
}
