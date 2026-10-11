package technitium

import "github.com/funtimecoding/soil/pkg/technitium/constant"

func WidenAccess(mode string) string {
	switch mode {
	case constant.AccessDeny:
		return constant.AccessNetwork
	case constant.AccessNameServer:
		return constant.AccessNameServerAndNetwork
	default:
		return mode
	}
}
