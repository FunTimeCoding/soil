package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func CustomFields(v []*custom_field.Field) []*server.CustomField {
	result := make([]*server.CustomField, 0, len(v))

	for _, f := range v {
		result = append(result, CustomField(f))
	}

	return result
}
