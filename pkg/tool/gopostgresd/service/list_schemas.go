package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
)

func (s *Service) ListSchemas(
	x context.Context,
	instance string,
) (*result.Result, error) {
	return s.guarded(x, instance, ours, constant.ListSchemasQuery)
}
