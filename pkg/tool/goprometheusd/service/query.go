package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"time"
)

func (s *Service) Query(
	x context.Context,
	instance string,
	query string,
	t time.Time,
) (*query_result.Result, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.Client) (*query_result.Result, error) {
			return c.Query(query, t)
		},
	)
}
