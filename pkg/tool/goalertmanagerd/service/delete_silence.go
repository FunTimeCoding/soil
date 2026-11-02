package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
)

func (s *Service) DeleteSilence(
	x context.Context,
	instance string,
	identifier string,
) error {
	_, e := guarded(
		s,
		x,
		instance,
		func(c face.AlertmanagerClient) (struct{}, error) {
			return struct{}{}, c.DeleteSilence(identifier)
		},
	)

	return e
}
