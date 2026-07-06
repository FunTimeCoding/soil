package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
)

func (s *Service) Alerts(
	instance string,
	f *alert_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, nil, e
	}

	return c.Alerts(nil, f)
}
