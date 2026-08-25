package collector

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"

func backupLabel() []string {
	return []string{
		constant.HypervisorLabel,
		constant.TypeLabel,
		constant.IdentifierLabel,
		constant.NameLabel,
	}
}
