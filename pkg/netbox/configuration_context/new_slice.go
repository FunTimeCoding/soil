package configuration_context

import "github.com/funtimecoding/soil/pkg/netbox/configuration_record"

func NewSlice(v []configuration_record.Record) []*Context {
	var result []*Context

	for i := range v {
		result = append(result, New(&v[i]))
	}

	return result
}
