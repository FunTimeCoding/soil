package service

import "github.com/funtimecoding/soil/pkg/netbox/custom_field"

func (s *Service) CreateSelectionField(
	name string,
	label string,
	choiceSet string,
) (*custom_field.Field, error) {
	return s.client.CreateSelectionField(name, label, choiceSet)
}
