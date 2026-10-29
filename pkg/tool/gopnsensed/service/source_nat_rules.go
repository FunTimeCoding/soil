package service

import "github.com/funtimecoding/soil/pkg/opnsense/source_nat"

func (s *Service) SourceNatRules(phrase string) ([]*source_nat.Rule, error) {
	return s.source.SourceNatRules(phrase)
}
