package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func CustomField(v *custom_field.Field) *server.CustomField {
	result := &server.CustomField{
		Identifier:  v.Identifier,
		Name:        v.Name,
		Kind:        v.Kind,
		ObjectTypes: v.ObjectTypes,
	}

	if result.ObjectTypes == nil {
		result.ObjectTypes = []string{}
	}

	if v.Label != "" {
		result.Label = &v.Label
	}

	if v.ChoiceSet != "" {
		result.ChoiceSet = &v.ChoiceSet
	}

	return result
}
