package service

import "github.com/funtimecoding/soil/pkg/opnsense/source_nat"

func (s *Service) SourceNatRules(phrase string) ([]*source_nat.Rule, error) {
	return guarded(
		s.guard,
		func() ([]*source_nat.Rule, error) {
			return s.client.SourceNatRules(phrase)
		},
	)
}
