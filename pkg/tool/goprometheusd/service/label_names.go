package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (s *Service) LabelNames(
	x context.Context,
	instance string,
	matches []string,
	span time.Duration,
) (*label_result.Result, error) {
	now := s.clock()
	start := spanStart(now, span)

	return guarded(
		s,
		x,
		instance,
		func(c face.PrometheusClient) (*label_result.Result, error) {
			return c.LabelNames(matches, start, now)
		},
	)
}
