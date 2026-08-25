package collector

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"

func guestLabel() []string {
	return []string{
		constant.HypervisorLabel,
		constant.NodeLabel,
		constant.TypeLabel,
		constant.IdentifierLabel,
		constant.NameLabel,
	}
}
