package service

import "github.com/funtimecoding/soil/pkg/prometheus/rule"

func (s *Service) Rules(instance string) ([]*rule.Rule, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	v, f := c.Rules()

	if f != nil {
		return nil, f
	}

	return v.Get(), nil
}
