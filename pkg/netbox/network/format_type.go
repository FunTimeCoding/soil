package network

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func (i *Interface) formatType() string {
	result := string(i.Type)

	if result == "" {
		result = constant.NoType
	}

	return result
}
