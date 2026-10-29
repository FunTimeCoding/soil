package service

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (s *Service) ChoiceSetByName(n string) (*choice_set.Set, error) {
	return s.client.ChoiceSetByName(n)
}
