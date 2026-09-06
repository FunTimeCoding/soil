package inventory

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"

func New(instances ...instance.Instance) *Inventory {
	return &Inventory{Instances: instances}
}
