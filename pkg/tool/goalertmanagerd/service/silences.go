package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
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
		func(c face.Client) ([]*silence.Silence, error) {
			return c.Silences(expired)
		},
	)
}
