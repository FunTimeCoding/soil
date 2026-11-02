package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
)

func (s *Service) Alerts(
	x context.Context,
	instance string,
	f *alert_filter.Filter,
) ([]*alert.Alert, *statistic.Statistic, error) {
	r, e := s.rules(x, instance)

	if e != nil {
		return nil, nil, e
	}

	p, e := s.pool.Client(instance)

	if e != nil {
		return nil, nil, e
	}

	o := advanced_option.New()
	o.All = true
	var v []*alert.Alert
	var t *statistic.Statistic
	g := p.AlertmanagerGuard().Attempt(
		x,
		func(context.Context) error {
			var h error
			v, t, h = p.Alertmanager().Alerts(o, f, r)

			return h
		},
	)

	return v, t, g
}
