package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"time"
)

func (s *Service) CreateSilence(
	x context.Context,
	instance string,
	alert string,
	comment string,
	d time.Duration,
) (string, error) {
	r, e := s.rules(x, instance)

	if e != nil {
		return "", e
	}

	return guarded(
		s,
		x,
		instance,
		func(c face.AlertmanagerClient) (string, error) {
			return c.SetSilence(alert, comment, d, r)
		},
	)
}
