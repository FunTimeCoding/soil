package choice_set

import "github.com/funtimecoding/soil/pkg/netbox/choice_set_record"

func NewSlice(v []choice_set_record.Record) []*Set {
	var result []*Set

	for i := range v {
		result = append(result, New(&v[i]))
	}

	return result
}
