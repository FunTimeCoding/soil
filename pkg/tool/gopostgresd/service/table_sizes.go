package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
)

func (s *Service) TableSizes(
	x context.Context,
	instance string,
	schema string,
) (*result.Result, error) {
	return s.guarded(
		x,
		instance,
		ours,
		constant.TableSizesQuery,
		schemaOrDefault(schema),
	)
}
