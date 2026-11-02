package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func Inventory() *instance.Inventory {
	return containerOption().Inventory
}
