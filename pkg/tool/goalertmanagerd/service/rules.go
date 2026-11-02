package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
)

func (s *Service) rules(
	x context.Context,
	instance string,
) (*rule_list.List, error) {
	p, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	var result *rule_list.List
	f := p.PrometheusGuard().Attempt(
		x,
		func(context.Context) error {
			var g error
			result, g = p.Prometheus().Rules()

			return g
		},
	)

	return result, f
}
