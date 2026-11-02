package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (s *Service) Status(
	x context.Context,
	instance string,
) (*models.AlertmanagerStatus, error) {
	return guarded(
		s,
		x,
		instance,
		func(c face.Client) (*models.AlertmanagerStatus, error) {
			return c.Status()
		},
	)
}
