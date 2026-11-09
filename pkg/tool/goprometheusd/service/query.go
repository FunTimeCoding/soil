package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
)

func (s *Service) Query(
	x context.Context,
	instance string,
	query string,
) (*query_result.Result, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.PrometheusClient) (*query_result.Result, error) {
			return c.Query(query, s.clock())
		},
	)
}
