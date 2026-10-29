package service

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (s *Service) SetChoiceSetChoices(
	name string,
	choices []string,
) (*choice_set.Set, error) {
	return s.client.SetChoiceSetChoices(name, choices)
}
