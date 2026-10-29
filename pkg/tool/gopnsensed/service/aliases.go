package service

import "github.com/funtimecoding/soil/pkg/opnsense/alias"

func (s *Service) Aliases(phrase string) ([]*alias.Alias, error) {
	return s.source.Aliases(phrase)
}
