package collector

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"

func storageLabel() []string {
	return []string{
		constant.HypervisorLabel,
		constant.NodeLabel,
		constant.StorageLabel,
		constant.PluginLabel,
		constant.ContentLabel,
	}
}
