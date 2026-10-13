package option

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

type Proxmox struct {
	Address       string
	ServiceTokens []string
	MetricAddress string
	Inventory     *instance.Inventory
}
