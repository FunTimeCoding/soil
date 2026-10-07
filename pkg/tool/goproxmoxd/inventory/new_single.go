package inventory

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"

func NewSingle(name string) *Inventory {
	return &Inventory{
		Instances: []instance.Instance{{Name: name, Host: "mock"}},
	}
}
