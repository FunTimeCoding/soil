package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (s *Service) QueryRange(
	x context.Context,
	instance string,
	query string,
	r v1.Range,
) (*query_result.Result, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.Client) (*query_result.Result, error) {
			return c.QueryRange(query, r)
		},
	)
}
