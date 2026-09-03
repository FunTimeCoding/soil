package option

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"

type Proxmox struct {
	Address       string
	ServiceTokens []string
	MetricAddress string
	Inventory     *inventory.Inventory
	Version       string
}
