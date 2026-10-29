package service

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (s *Service) CreateChoiceSet(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	return s.client.CreateChoiceSet(name, choices)
}
