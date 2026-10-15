package choice_set

import "github.com/funtimecoding/soil/pkg/netbox/choice_set_record"

func New(v *choice_set_record.Record) *Set {
	result := &Set{
		Identifier: v.Identifier,
		Name:       v.Name,
		Choices:    []string{},
		Raw:        v,
	}

	for _, c := range v.Choices {
		if len(c) > 0 {
			result.Choices = append(result.Choices, c[0])
		}
	}

	return result
}
