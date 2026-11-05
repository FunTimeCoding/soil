package service

import "github.com/funtimecoding/soil/pkg/opnsense/rule"

func (s *Service) Rules(phrase string) ([]*rule.Rule, error) {
	return guarded(
		s.guard,
		func() ([]*rule.Rule, error) {
			return s.client.Rules(phrase)
		},
	)
}
