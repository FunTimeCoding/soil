package service

import "github.com/funtimecoding/soil/pkg/netbox/custom_field"

func (s *Service) CustomFields() ([]*custom_field.Field, error) {
	return s.client.CustomFields()
}
