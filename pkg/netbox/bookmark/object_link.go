package bookmark

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/helper"
)

func objectLink(object any) string {
	fields, k := object.(map[string]any)

	if !k {
		return ""
	}

	raw, l := fields[constant.ObjectLinkField].(string)

	if !l {
		return ""
	}

	return helper.ToWebLink(raw)
}
