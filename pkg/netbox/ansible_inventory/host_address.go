package ansible_inventory

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"strings"
)

func hostAddress(primary string) string {
	if primary == constant.NoPrimaryAddress {
		return ""
	}

	result, _, _ := strings.Cut(primary, "/")

	return result
}
