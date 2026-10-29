package service

import "github.com/funtimecoding/soil/pkg/netbox/choice_set"

func (s *Service) ChoiceSets() ([]*choice_set.Set, error) {
	return s.client.ChoiceSets()
}
