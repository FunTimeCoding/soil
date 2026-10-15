package configuration_context

import "github.com/funtimecoding/soil/pkg/netbox/configuration_record"

func names(v []configuration_record.Named) []string {
	result := []string{}

	for _, n := range v {
		result = append(result, n.Name)
	}

	return result
}
