package object_type

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func Label(v string) string {
	if result, k := constant.ObjectTypeAlias[v]; k {
		return result
	}

	return v
}
