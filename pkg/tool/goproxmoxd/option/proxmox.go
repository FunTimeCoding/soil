package option

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/inventory"

type Proxmox struct {
	Port      int
	Inventory *inventory.Inventory
	Version   string
}
