package custom_field

import "github.com/funtimecoding/soil/pkg/netbox/custom_field_record"

func New(v *custom_field_record.Record) *Field {
	result := &Field{
		Identifier:  v.Identifier,
		Name:        v.Name,
		Label:       v.Label,
		Kind:        v.Kind.Value,
		ObjectTypes: v.ObjectTypes,
		Raw:         v,
	}

	if v.ChoiceSet != nil {
		result.ChoiceSet = v.ChoiceSet.Name
	}

	return result
}
