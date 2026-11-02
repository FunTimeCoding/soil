package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

func (s *Service) Silences(
	x context.Context,
	instance string,
	expired bool,
) ([]*silence.Silence, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.AlertmanagerClient) ([]*silence.Silence, error) {
			return c.Silences(expired)
		},
	)
}
