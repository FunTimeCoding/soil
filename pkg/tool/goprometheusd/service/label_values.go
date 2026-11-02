package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (s *Service) LabelValues(
	x context.Context,
	instance string,
	label string,
	matches []string,
	since time.Time,
) (*label_result.Result, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.Client) (*label_result.Result, error) {
			return c.LabelValues(label, matches, since)
		},
	)
}
