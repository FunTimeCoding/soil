package service

import "context"

func (s *Service) ListSchemas(
	x context.Context,
	instance string,
) ([]map[string]any, error) {
	return s.Query(x, instance, listSchemasSQL)
}
