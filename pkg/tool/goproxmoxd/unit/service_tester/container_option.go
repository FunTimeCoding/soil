package service_tester

import (
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func containerOption() *option.Proxmox {
	result := option.New()
	result.Inventory = inventory.New[instance.Instance](
		instance.Instance{Name: "test", Host: "mock"},
	)

	return result
}
