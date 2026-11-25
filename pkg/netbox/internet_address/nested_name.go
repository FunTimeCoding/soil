package internet_address

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func nestedName(
	object map[string]any,
	key string,
) string {
	nested, okay := object[key].(map[string]any)

	if !okay {
		return ""
	}

	return text(nested, constant.NameField)
}
