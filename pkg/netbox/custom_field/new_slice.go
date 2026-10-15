package custom_field

import "github.com/funtimecoding/soil/pkg/netbox/custom_field_record"

func NewSlice(v []custom_field_record.Record) []*Field {
	var result []*Field

	for i := range v {
		result = append(result, New(&v[i]))
	}

	return result
}
