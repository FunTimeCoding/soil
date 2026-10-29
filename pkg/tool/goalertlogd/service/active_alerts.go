package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
)

func (s *Service) activeAlerts() ([]*alert.Alert, error) {
	var result []*alert.Alert
	e := s.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			var f error
			result, _, f = s.alertmanager.Alerts(advanced_option.New(), nil)

			return f
		},
	)

	return result, e
}
