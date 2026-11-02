package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (s *Service) Series(
	x context.Context,
	instance string,
) (v1.TSDBResult, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.Client) (v1.TSDBResult, error) {
			return c.Series()
		},
	)
}
