package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
)

func (s *Service) ListSchemas(
	x context.Context,
	instance string,
) ([]map[string]any, error) {
	return s.Query(x, instance, constant.ListSchemasQuery)
}
