package inventory

import "github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"

func New(instances ...instance.Instance) *Inventory {
	return &Inventory{Instances: instances}
}
